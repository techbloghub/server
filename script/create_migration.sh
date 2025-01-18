#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <migration_name>"
  exit 1
fi

MIGRATION_NAME=$1

atlas migrate diff "$MIGRATION_NAME" \
  --dir "file://ent/migrate/migrations" \
  --to "ent://ent/schema" \
  --dev-url "docker://postgres/15/test?search_path=public"
