-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;


-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
-- Option one with flags and case statement
-- UPDATE users
-- SET
--   hashed_password = CASE
--     WHEN @set_hashed_password::boolean = TRUE THEN @hashed_password::text
--     ELSE hashed_password
--   END,
--   full_name = CASE
--     WHEN @set_full_name::boolean = TRUE THEN @full_name::text
--     ELSE full_name
--   END,
--   email = CASE
--     WHEN @set_email::boolean = TRUE THEN @email::text
--     ELSE email
--   END
-- WHERE
--   username = @username::text
-- RETURNING *;

-- Better option using nullable fields with coalesce
UPDATE users
SET
  hashed_password = coalesce(sqlc.narg(hashed_password), hashed_password),
  password_changed_at = coalesce(sqlc.narg(password_changed_at), password_changed_at),
  full_name = coalesce(sqlc.narg(full_name), full_name),
  email = coalesce(sqlc.narg(email), email)
WHERE
  username = sqlc.arg(username)
RETURNING *;
