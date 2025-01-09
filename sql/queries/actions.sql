-- ----------CREATE Section - START-------------

-- Create Action entity
-- name: CreateAction :one
INSERT INTO actions (id, created_at, updated_at, name, description, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- ----------CREATE Section - END-------------//