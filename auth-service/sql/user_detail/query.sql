-- name: CreateUserDetail :exec
INSERT INTO user_details (
    user_guid, 
    first_name, 
    last_name, 
    phone, 
    avatar, 
    address, 
    createdAt
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetUserDetailByUserGuid :one
SELECT * FROM user_details 
WHERE deletedAt IS NULL AND user_guid = ?;

-- name: UpdateUserDetailByUserGuid :exec
UPDATE user_details 
SET 
    first_name = COALESCE(sqlc.narg('first_name'), first_name),
    last_name = COALESCE(sqlc.narg('last_name'), last_name),
    phone = COALESCE(sqlc.narg('phone'), phone),
    avatar = COALESCE(sqlc.narg('avatar'), avatar),
    address = COALESCE(sqlc.narg('address'), address),
    updatedAt = sqlc.arg('updatedAt')
WHERE user_guid = sqlc.arg('user_guid');

-- name: DeleteUserDetailByUserGuid :exec
UPDATE user_details 
SET deletedAt = ? 
WHERE user_guid = ?;