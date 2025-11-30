-- +goose Up
ALTER TABLE repos ADD COLUMN state integer NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE repos DROP COLUMN state;
