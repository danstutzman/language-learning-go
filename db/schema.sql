DROP TABLE IF EXISTS cards;
CREATE TABLE cards (
  id                INTEGER PRIMARY KEY NOT NULL,
  l1                TEXT NOT NULL,
  l2                TEXT NOT NULL
);
CREATE INDEX idx_cards_l1 ON cards(l1);
CREATE INDEX idx_cards_l2 ON cards(l2);

DROP TABLE IF EXISTS cards_morphemes;
CREATE TABLE cards_morphemes (
  card_id INTEGER NOT NULL,
  morpheme_id INTEGER NOT NULL,
  num_morpheme INTEGER NOT NULL
);
CREATE INDEX idx_cards_morphemes_card_id ON cards_morphemes(card_id);
CREATE INDEX idx_cards_morphemes_morpheme_id ON cards_morphemes(morpheme_id);

DROP TABLE IF EXISTS morphemes;
CREATE TABLE morphemes (
  id                INTEGER PRIMARY KEY NOT NULL,
  type              TEXT NOT NULL,
  l2                TEXT NOT NULL,
  lemma             TEXT, -- type=VERB_SUFFIX has no lemma
  freeling_tag      TEXT,
  verb_category     TEXT
);
CREATE UNIQUE INDEX idx_morphemes_type_and_l2 ON morphemes(type, l2);
CREATE UNIQUE INDEX idx_morphemes_verb_category_and_freeling_tag
  ON morphemes(verb_category, freeling_tag);
