#!/usr/bin/env ruby

# echo "select * from morphemes where type = 'VERB_STEM_CHANGE'" | sqlite3 db/db.sqlite3 > db/verb_stem_changes.csv

File.open('db/verb_stem_changes.csv').each_line do |line|
  _, _, stem_change, infinitive = line.strip.split('|')
  if infinitive.end_with?('ar')
    infinitive.gsub! /ar$/, ''
    if infinitive.gsub(/z$/, 'c-') == stem_change
      #STDERR.puts "Explained #{stem_change}"
    elsif infinitive.gsub(/c$/, 'qu-') == stem_change
      #STDERR.puts "Explained #{stem_change}"
    elsif infinitive.gsub(/o([bcdfgjklmnpqrstvwzñ]+)$/, 'ue\1-') == stem_change
      #STDERR.puts "Explained #{stem_change}"
    elsif infinitive.gsub(/e([bcdfgjklmnpqrstvwzñ]+)$/, 'ie\1-') == stem_change
      #STDERR.puts "Explained #{stem_change}"
    else
      STDERR.puts "Can't explain #{stem_change} for #{infinitive}"
    end
  end
end
