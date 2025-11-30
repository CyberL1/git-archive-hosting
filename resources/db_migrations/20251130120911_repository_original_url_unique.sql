-- +goose Up
CREATE UNIQUE INDEX IF NOT EXISTS repos_original_url_key ON repos (original_url);

-- +goose Down
DROP INDEX IF EXISTS repos_original_url_key;
