#!/usr/bin/env ruby
require 'sqlite3'

ONSETS = %w[b C d D f g h J k l m n p r s S t T v w y z Z] # removed N
RIMES = %w[AE EY AO AX IY EH IH AY IX AA UW UH UX OW AW OY]

db = SQLite3::Database.new("#{__dir__}/../db/db.sqlite3")

nonsense_by_morpheme_id = {}
db.execute('select id, l2 from morphemes') do |row|
  id = row[0]
  pattern = row[1]
    .downcase
    .gsub(/[aeiouáéíóú]+/, 'V')
    .gsub(/[bcdfghjklmnpqrstvwyz]+/, 'C')

  nonsense = ''
  pattern.each_char do |char|
    if char == 'C'
      nonsense += ' ' + ONSETS[rand(ONSETS.size)]
    elsif char == 'V'
      nonsense += RIMES[rand(RIMES.size)]
    else
      nonsense += char # punctuation
    end
  end

  nonsense_by_morpheme_id[id] = nonsense.strip
end

#begin
#  db.execute("alter table morphemes add column nonsense TEXT")
#rescue SQLite3::SQLException => e
#  raise unless e.message == 'duplicate column name: nonsense'
#end

for morpheme_id, nonsense in nonsense_by_morpheme_id
  db.execute("update morphemes set nonsense='#{nonsense.gsub("'", "''")}'
             where id = #{morpheme_id}")
end
