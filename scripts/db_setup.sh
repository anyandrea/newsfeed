#!/bin/bash

# fail on error
set -e

# =============================================================================================
if [[ "$(basename $PWD)" == "scripts" ]]; then
	cd ..
fi
echo $PWD

# =============================================================================================
source .env_mysql

# =============================================================================================
echo "setting up newsfeeddb ..."

while [[ ! mysqladmin ping -h"127.0.0.1" --silent ]]; do
	echo "waiting for mysql to be ready ..."
    sleep 2
done

mysql --host=127.0.0.1 --user=blubb --password=blabb --database=newsfeed_db -v < lib/database/migrations/mysql/001_newsfeeddb_user_tables.up.sql
mysql --host=127.0.0.1 --user=blubb --password=blabb --database=newsfeed_db -v < lib/database/migrations/mysql/002_newsfeeddb_feed_tables.up.sql
mysql --host=127.0.0.1 --user=blubb --password=blabb --database=newsfeed_db -v < _fixtures/users.sql
mysql --host=127.0.0.1 --user=blubb --password=blabb --database=newsfeed_db -v < _fixtures/feeds.sql
