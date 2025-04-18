
-- name: GetUserById :one
SELECT * FROM users WHERE deletedAt IS NULL AND id = ?;

-- name: GetUserByGuid :one
SELECT * FROM users WHERE deletedAt IS NULL AND guid = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE deletedAt IS NULL AND siteId = ? AND email = ? ;

-- name: GetUsersPaging :many
SELECT * FROM users WHERE deletedAt IS NULL AND siteId = ? ORDER BY createdAt DESC LIMIT ? OFFSET ?;


-- name: CreateUser :exec
INSERT INTO users (guid, siteId, email, hash_password, salt, createdAt) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateUserById :exec
UPDATE users SET siteId = ?, email = ?, hash_password = ?, salt = ?, updatedAt = ? WHERE id = ?;

-- name: UpdateUserByGuid :exec
UPDATE users SET siteId = ?, email = ?, hash_password = ?, salt = ?, updatedAt = ? WHERE guid = ?;

-- name: DeleteUserById :exec
UPDATE users SET deletedAt = ? WHERE id = ?;

-- name: DeleteUserByGuid :exec
UPDATE users SET deletedAt = ? WHERE guid = ?;