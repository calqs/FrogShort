#!/bin/bash
set -e

echo "Waiting for database connection..."
until pg_isready; do
  sleep 2
done
echo "Database is ready."

# Check if schema_migrations table exists
TABLE_EXISTS=$(psql -tAc "SELECT to_regclass('${DB_SCHEMA}.schema_migrations');")
echo "schema_migrations exists? [$TABLE_EXISTS]"
if [ -z "$TABLE_EXISTS" ]; then
  echo "No migration table found, applying initial SQL..."
  echo "Applying /migrations/init.sql..."
  psql -v ON_ERROR_STOP=1 -f /migrations/init.sql
  echo "Creating migration tracking table..."
  psql -v ON_ERROR_STOP=1 -f /migrations/create_migrations_table.sql
  echo "Initial migration applied successfully."
else
  echo "schema_migrations table already exists, skipping init scripts."
fi

echo "Running 'up' migrations..."

# Loop over all .sql files in /migrations/up, in lexicographic order
for file in /migrations/up/*.sql; do
  # if no files match, the glob is returned literally; skip in that case
  if [ ! -e "$file" ]; then
    echo "No migration files found in /migrations/up, nothing to do."
    break
  fi

  filename=$(basename "$file")

  # Check if this migration was already applied
  ALREADY_APPLIED=$(psql -tAc \
    "SELECT 1 FROM ${DB_SCHEMA}.schema_migrations WHERE filename = '$filename' LIMIT 1;")

  if [ -z "$ALREADY_APPLIED" ]; then
    echo "Applying migration: $filename"
    psql -v ON_ERROR_STOP=1 -f "$file"

    echo "Recording migration: $filename"
    psql -v ON_ERROR_STOP=1 -c \
      "INSERT INTO ${DB_SCHEMA}.schema_migrations (filename) VALUES ('$filename');"
  else
    echo "Skipping already applied migration: $filename"
  fi
done

echo "All migrations processed."
