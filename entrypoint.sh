#!/usr/bin/env bash
# Creates .pgpass with connection data from env variables for pg_dump.
# 1. pg_dump asks for password if .pgpass does not exist
# 2. It's easier to set up containers in one file "docker-compose.yml" via env variables.
set -e

echo "${PG_HOST}":"${PG_PORT}":"${PG_DB}":"${PG_USER}":"${PG_PASSWORD}" > ~/.pgpass
chmod 0600 ~/.pgpass

exec "$@"