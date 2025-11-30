-- name: ListRepos :many
SELECT * FROM repos;

-- name: ListReposBySource :many
SELECT * FROM repos WHERE LOWER(source) = LOWER(sqlc.arg(source));

-- name: ListReposBySourceAndOwner :many
SELECT * FROM repos WHERE LOWER(source) = LOWER(sqlc.arg(source)) AND LOWER(owner) = LOWER(sqlc.arg(owner));

-- name: GetRepoById :one
SELECT * FROM repos WHERE id = ?;

-- name: GetRepoByFullName :one
SELECT * FROM repos WHERE LOWER(source) = LOWER(sqlc.arg(source)) AND LOWER(owner) = LOWER(sqlc.arg(owner)) AND LOWER(name) = LOWER(sqlc.arg(name));

-- name: CreateRepo :one
INSERT INTO repos (owner, name, original_url, created_at, source, state) VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateRepoState :one
UPDATE repos SET state = ? WHERE id = ?
RETURNING *;
