#!/bin/bash
set -x
if [ -f .env ]; then
    source .env
fi

cd sql/schema
echo "Migrating database with URL: $DATABASE_URL"
goose turso $DATABASE_URL up
