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
echo "connecting to newsfeeddb ..."

mysql --host=127.0.0.1 --user=blubb --password=blabb --database=newsfeed_db
