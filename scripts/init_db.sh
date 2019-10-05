#!/bin/bash -ex

cd `dirname $0`/..

rm -f db.sqlite3

sqlite3 db.sqlite3 <<EOF
  CREATE TABLE cards (
    id                INTEGER NOT NULL,
    l1                TEXT NOT NULL,
    l2                TEXT NOT NULL,
    morpheme_ids_json TEXT NOT NULL,
    created_at_millis REAL NOT NULL,
    updated_at_millis REAL NOT NULL
  );
  INSERT INTO cards(id, l1, l2, morpheme_ids_json, created_at_millis, updated_at_millis)
    VALUES(-1, 'man', 'hombre', '[]', STRFTIME('%s','now'), STRFTIME('%s','now'));
EOF

echo '.schema' | sqlite3 db.sqlite3
