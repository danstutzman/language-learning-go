require 'csv'

# Part 1: load verb suffixes and dump them

puts "BEGIN TRANSACTION;"

verb_suffixes = CSV.read("#{__dir__}/../db/2_verb_suffixes.csv", headers: true)

for verb_suffix in verb_suffixes
  puts sprintf(
    "INSERT INTO morphemes (type, l2, lemma, freeling_tag, verb_category)
    VALUES ('VERB_SUFFIX', '-%s', NULL, '%s', '%s');",
    verb_suffix['suffix'].gsub("'", "''"),
    verb_suffix['tag'].gsub("'", "''"),
    verb_suffix['category'].gsub("'", "''"))
end

# Part 2: load lemma2tag2form

category_tag2suffixes = {}
lemma2tag2form = {}

path = "#{__dir__}/../db/freeling/dicc.src"
File.open(path).each_with_index do |line, line_num0|
  line_num = line_num0 + 1
  next if line_num < 5

  values = line.strip.split(' ')
  form = values.shift

  while values.size > 0
    lemma = values.shift
    tag = values.shift

    if lemma.match(/[aei]r$/) && tag.start_with?('V')
      expected_stem = lemma[0...-2]
      if form.start_with?(expected_stem)
        category = lemma[-2..-1]
        found_suffix = false
        for verb_suffix in verb_suffixes
          if verb_suffix['category'] == category &&
              verb_suffix['tag'] == tag &&
              form == expected_stem + verb_suffix['suffix']
            found_suffix = true
            break
          end
        end
        if !found_suffix
          expected_suffix = form[expected_stem.size..-1]
          #STDERR.puts "No suffix found for #{form} lemma=#{lemma} tag=#{tag}"
          category_tag2suffixes[[category, tag]] ||= {}
          category_tag2suffixes[[category, tag]][expected_suffix] ||= 0
          category_tag2suffixes[[category, tag]][expected_suffix] += 1
        end
      else
        lemma2tag2form[lemma] ||= {}
        lemma2tag2form[lemma][tag] = form
      end
    end
  end
end

# Part 3: deduce what the stems are

# estar isn't caught by this script because estuv- matches est-, the
# expected stem of estar.  So we have to insert it manually.
puts "INSERT INTO morphemes
  (type, l2, lemma, freeling_tag, verb_category)
  VALUES ('VERB_STEM_CHANGE', 'estuv-', 'estar', NULL, NULL);"

for lemma, tag2form in lemma2tag2form
  new_stem2votes = Hash.new(0)
  unexplained = {}

  for tag, form in tag2form
    category = lemma[-2..-1]
    for verb_suffix in verb_suffixes
      if verb_suffix['category'] == category && verb_suffix['tag'] == tag &&
         form.end_with?(verb_suffix['suffix'])
        stem_length = form.size - verb_suffix['suffix'].size
        new_stem2votes[form[0...stem_length]] += 1
        break
      end
    end
  end

  for new_stem, votes in new_stem2votes
    if votes >= 3
      puts sprintf("INSERT INTO morphemes
        (type, l2, lemma, freeling_tag, verb_category)
        VALUES ('VERB_STEM_CHANGE', '%s', '%s', NULL, NULL);",
        new_stem.gsub("'", "''") + '-', lemma.gsub("'", "''"))
    else
      STDERR.puts "Skipped inserting #{lemma} -> #{new_stem}"
    end
  end
end

puts 'COMMIT TRANSACTION;'

# Uncomment to get a list of regular (not stem changing) suffixes
#for category_tag, suffixes in category_tag2suffixes
#  category, tag = category_tag
#  for suffix, count in suffixes
#    STDERR.puts [category, tag, suffix, count].inspect if count >= 100
#  end
#end
