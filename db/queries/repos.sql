-- name: ListRepos :many
SELECT * FROM repos;

-- name: GetRepoById :one
SELECT * FROM repos WHERE id = ?;

-- name: GetRepoByOwner :one
SELECT * FROM repos WHERE LOWER(owner) = LOWER(sqlc.arg(owner));

-- name: GetRepoByFullName :one
SELECT * FROM repos WHERE LOWER(owner) = LOWER(sqlc.arg(owner)) AND LOWER(name) = LOWER(sqlc.arg(name));

-- name: CreateRepo :one
INSERT INTO repos (owner, name, original_url, created_at) VALUES (?, ?, ?, ?)
RETURNING *;
