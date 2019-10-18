#!/usr/bin/env python3

# To install:
#   pip3 install spacy
#   python3 -m spacy download es_core_news_sm

import csv
import spacy
import sys

nlp = spacy.load('es_core_news_sm')

writer = csv.writer(sys.stdout)
writer.writerow(['line_num', 'index', 'text', 'norm', 'tag', 'lemma',
  'pos', 'dep', 'parent_index'])

line_num = 0
for line in sys.stdin:
  line_num += 1
  phrase = line.rstrip()
  doc = nlp(phrase)

  for token in doc:
    writer.writerow([line_num, token.i, token.text, token.norm_, token.tag_,
        token.lemma_, token.pos_, token.dep_, token.head.i])
