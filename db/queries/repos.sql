-- name: ListRepos :many
SELECT * FROM repos;

-- name: GetRepoById :one
SELECT * FROM repos WHERE id = ?;

-- name: GetRepoByOwner :one
SELECT * FROM repos WHERE owner = ?;

-- name: GetRepoByFullName :one
SELECT * FROM repos WHERE owner = ? AND name = ?;

-- name: CreateRepo :one
INSERT INTO repos (owner, name, original_url, created_at) VALUES (?, ?, ?, ?)
RETURNING *;
