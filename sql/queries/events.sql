----- CREATE - START ----

-- name: CreateEvent :one
INSERT into action_events(id, created_at, updated_at, executed_at, action_id, user_id, comment) 
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;


----- CREATE - END ----
