#!/bin/bash -ex

cd `dirname $0`/..

rm -f db.sqlite3

ruby ~/dev/language-learning-corpora/es-en-dictionary-from-anki-deck/convert_to_sql.rb | sqlite3 db.sqlite3

sqlite3 db.sqlite3 <<EOF
  CREATE TABLE cards (
    id                INTEGER PRIMARY KEY NOT NULL,
    l1                TEXT NOT NULL,
    l2                TEXT NOT NULL,
    morpheme_ids_json TEXT NOT NULL,
    created_at_millis REAL NOT NULL,
    updated_at_millis REAL NOT NULL
  );
  CREATE UNIQUE INDEX idx_cards_l1 ON cards(l1);
  CREATE UNIQUE INDEX idx_cards_l2 ON cards(l2);
  INSERT INTO cards (l1, l2, morpheme_ids_json, created_at_millis, updated_at_millis)
    VALUES ('man', 'hombre', '[]', 0, 0);

  CREATE TABLE cards_morphemes (
    card_id INTEGER NOT NULL,
    morpheme_id INTEGER NOT NULL
  );
  CREATE INDEX idx_cards_morphemes_card_id ON cards_morphemes(card_id);
  CREATE INDEX idx_cards_morphemes_morpheme_id ON cards_morphemes(morpheme_id);
  INSERT INTO cards_morphemes (card_id, morpheme_id) VALUES (
   (SELECT id FROM cards WHERE l2 = 'hombre'),
   (SELECT id FROM morphemes WHERE l2 = 'hombre'));
EOF

echo '.schema' | sqlite3 db.sqlite3
