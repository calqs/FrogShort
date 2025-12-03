-- migrations/init.sql
CREATE SCHEMA IF NOT EXISTS frogshort;
SET search_path TO frogshort;

CREATE EXTENSION IF NOT EXISTS "pg_uuidv7" WITH SCHEMA frogshort;

CREATE TABLE IF NOT EXISTS urls (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
  code TEXT UNIQUE NOT NULL,
  long_url TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);