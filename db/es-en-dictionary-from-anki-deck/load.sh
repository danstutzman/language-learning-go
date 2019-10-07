#!/bin/bash -ex
cd `dirname $0`
bundle exec ruby convert_to_sql.rb | sqlite3 -bail ../db.sqlite3
