#!/usr/bin/env ruby

if ARGV[0].nil?
  raise "1st arg: line num to play"
end
specified_line_num = ARGV[0].to_i

path = "#{__dir__}/../db/corpora/spanish-in-texas.txt"
File.open(path).each_line.with_index do |line, line_num|
  if line_num == specified_line_num
    line.strip!
    line.gsub! /^>>[is]: /, ''
    system("say -v Angelica '#{line}' -r 100")
    break
  end
end
