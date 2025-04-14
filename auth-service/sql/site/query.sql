-- name: CreateSite :exec
INSERT INTO sites (guid, siteId, name, createdAt) VALUES (?, ?, ?, ?);

-- name: GetSiteById :one
SELECT * FROM sites WHERE id = ?;

-- name: GetSiteByGuid :one
SELECT * FROM sites WHERE guid = ?;

-- name: UpdateSiteById :exec
UPDATE sites SET siteId = ?, name = ?, updatedAt = ? WHERE id = ?;

-- name: UpdateSiteByGuid :exec
UPDATE sites SET siteId = ?, name = ?, updatedAt = ? WHERE guid = ?;

-- name: DeleteSiteById :exec
UPDATE sites SET deletedAt = ? WHERE id = ?;

-- name: DeleteSiteByGuid :exec
UPDATE sites SET deletedAt = ? WHERE guid = ?;

-- name: GetSitesPaging :many
SELECT * FROM sites WHERE deletedAt IS NULL ORDER BY createdAt DESC LIMIT ? OFFSET ?;

SELECT * FROM sites WHERE deletedAt IS NULL ORDER BY createdAt DESC LIMIT ? OFFSET ?;

-- name: CountSites :one
SELECT COUNT(*) AS total FROM sites WHERE deletedAt IS NULL;
