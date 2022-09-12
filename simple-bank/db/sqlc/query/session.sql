-- name: CreateSession :one
INSERT INTO "sessions" (
  id, -- this should always be the refresh token id... not sure what I think about that.
  username,
  refresh_token,
  user_agent,
  client_ip,
  is_blocked,
  expires_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;


-- name: GetSession :one
SELECT * FROM "sessions"
WHERE id = $1 LIMIT 1;
