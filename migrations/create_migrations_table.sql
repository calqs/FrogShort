-- migrations/init.sql
SET search_path TO spendbaker;

CREATE TABLE schema_migrations (
    id UUID PRIMARY KEY uuid_generate_v7(),
    filename varchar(255) NOT NULL UNIQUE,
    applied_at TIMESTAMP DEFAULT now()
);