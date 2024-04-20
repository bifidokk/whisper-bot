#!/bin/bash
set -e
export PGPASSWORD=postgres;
psql -v ON_ERROR_STOP=1 --username "postgres" <<-EOSQL
  CREATE DATABASE whisperbot;
  GRANT ALL PRIVILEGES ON DATABASE whisperbot TO "postgres";
EOSQL