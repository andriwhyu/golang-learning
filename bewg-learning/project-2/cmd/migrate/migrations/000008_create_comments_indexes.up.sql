CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX IF NOT EXISTS comments_content_idx ON comments USING GIN (content gin_trgm_ops);