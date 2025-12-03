-- migrations/init.sql
CREATE SCHEMA IF NOT EXISTS frogshort;
SET search_path TO frogshort;

CREATE EXTENSION IF NOT EXISTS "pg_uuidv7" WITH SCHEMA frogshort;

CREATE TABLE IF NOT EXISTS urls (
  filename TEXT NOT NULL,
  applied_at TIMESTAMP NOT NULL DEFAULT NOW()
);