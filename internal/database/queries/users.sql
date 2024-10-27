-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

-- name: CreateUser :one
INSERT INTO users (
  id,
  username,
  created_at
) VALUES (
  gen_random_uuid(), -- PostgreSQL's built-in UUID generator
  $1,
  NOW()
) RETURNING *;
