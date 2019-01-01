#!/bin/bash -ex
cd `dirname $0`/..
rm -f db.sqlite3
sqlite3 db.sqlite3 <<EOF
  CREATE TABLE cards (
    cardId INTEGER PRIMARY KEY NOT NULL,
    es TEXT NOT NULL
  );
  CREATE TABLE exposures (
    createdAt REAL NOT NULL
  );
  INSERT INTO cards(es) VALUES('hombre');
  INSERT INTO cards(es) VALUES('mujer');
EOF
echo '.schema' | sqlite3 db.sqlite3
