#!/usr/bin/env ruby
require 'sqlite3'

# Onset N ("ng") was removed
ONSET_PHONES = %w[b C  d D  f g h J k l m n p r s S  t  T v w y z Z]
ONSET_L2S    = %w[b ch d dh f g h j k l m n p r s sh t th v w y z zh]

RIME_PHONES = %w[AE EY AO AX IY EH IH AY IX AA UW UH UX OW AW OY]
RIME_L2S    = %w[aa ei  a  a  i  e  i ai  i  a  u  u  a  o au oi]

L2_BY_NONSENSE_L2 = {}

db = SQLite3::Database.new("#{__dir__}/../db/db.sqlite3")

nonsense_l2_by_morpheme_id = {}
nonsense_phones_by_morpheme_id = {}
db.execute('select id, l2 from morphemes') do |row|
  id = row[0]
  l2 = row[1]
  pattern = l2
    .downcase
    .gsub(/[aeiouáéíóú]+/, 'V')
    .gsub(/[bcdfghjklmnpqrstvwyz]+/, 'C')

  need_assignment = true
  while need_assignment
    nonsense_l2 = ''
    nonsense_phones = ''
    pattern.each_char do |char|
      if char == 'C'
        onset_num = rand(ONSET_PHONES.size)
        nonsense_phones += '-' if nonsense_phones.size > 0
        nonsense_phones += ONSET_PHONES[onset_num]
        nonsense_l2 += ONSET_L2S[onset_num]
      elsif char == 'V'
        rime_num = rand(RIME_PHONES.size)
        nonsense_phones += RIME_PHONES[rime_num]
        nonsense_l2 += RIME_L2S[rime_num]
      else
        nonsense_l2 += char # punctuation
      end
    end

    # Regenerate it if there's already nonsense_l2 for a different meaning
    need_assignment = L2_BY_NONSENSE_L2[nonsense_l2] &&
      L2_BY_NONSENSE_L2[nonsense_l2] != l2
  end
  L2_BY_NONSENSE_L2[nonsense_l2] = l2

  nonsense_l2_by_morpheme_id[id] = nonsense_l2
  nonsense_phones_by_morpheme_id[id] = nonsense_phones
end

for morpheme_id, nonsense_l2 in nonsense_l2_by_morpheme_id
  nonsense_phones = nonsense_phones_by_morpheme_id[morpheme_id]
  db.execute("update morphemes set nonsense_l2='#{nonsense_l2.gsub("'", "''")}',
             nonsense_phones='#{nonsense_phones.gsub("'", "''")}'
             where id = #{morpheme_id}")
end
