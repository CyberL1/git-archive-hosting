-- +goose Up
CREATE TABLE repos (
  id integer PRIMARY KEY AUTOINCREMENT,
  owner text NOT NULL,
  name text NOT NULL,
  original_url text NOT NULL,
  created_at timestamp NOT NULL
);

-- +goose Down
DROP TABLE repos;
