-- name: CreateToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
        $1,
        NOW(),
        NOW(),
        $2,
        $3,
        NULL
       )
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT u.*
FROM refresh_tokens rt JOIN users u ON rt.user_id = u.id
WHERE rt.revoked_at IS NULL and rt.expires_at > NOW() and rt.token = $1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens SET revoked_at = NOW(), updated_at = NOW() WHERE token = $1;