require 'sqlite3'

db = SQLite3::Database.new("#{__dir__}/collection.anki2")

puts "BEGIN TRANSACTION;"

db.execute('select flds from notes') do |row|
  es_plus, _, english, _, _ = row[0].split("\u001F")
  next if es_plus.start_with?('Key to Abbreviations')

  es_plus.gsub! %r[.*?<br ?/?>], ''
  es_plus.gsub! %r[</?div>], ''
  english.gsub! %r[</?div>], ''
  english.gsub! %r[<br />$], ''
  english.gsub! %r[<img src=\"paste-81720342740993.jpg\" />$], ''
  english.gsub! %r[&nbsp;.*], ''

  match = /^(.*?)( |&nbsp;)\(([a-z,\/ ]+)\).*$/.match(es_plus)
  raise "Can't parse: #{es_plus}" if match.nil?
  es = match[1]
  part_of_speech = match[3]

  if part_of_speech == 'v' && match = /^(.*)[aei]r$/.match(es)
    l2 = match[1] + '-'
  else
    l2 = es
  end

  puts "INSERT INTO morphemes (gloss, l2, lemma, tags_csv) VALUES ('#{english.gsub("'", "''")}', '#{l2.gsub("'", "''")}', '#{es.gsub("'", "''")}', '#{part_of_speech.gsub("'", "''")}');"
end

puts "COMMIT TRANSACTION;"
