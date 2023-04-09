-- name: uuid-generate-extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- name: create-database
SELECT 'CREATE DATABASE vnexpress_crawler'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'vnexpress_crawler');