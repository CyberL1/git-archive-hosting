-- +goose Up
ALTER TABLE repos ADD COLUMN source text NOT NULL DEFAULT '';

UPDATE repos
SET source = LOWER(
  SUBSTR(
    original_url,
    INSTR(original_url, '://') + 3,
    INSTR(SUBSTR(original_url, INSTR(original_url, '://') + 3), '/') - 1
  )
);

-- +goose Down
ALTER TABLE repos DROP COLUMN source;
