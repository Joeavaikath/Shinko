-- ----------CREATE Section - START-------------

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, email, password_hash)
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


-- ----------RETRIEVE Section - START-------------
-- name: GetUserByEmail :one
SELECT * FROM users
WHERE
email = $1;
-- ----------RETRIEVE Section - END-------------//


-- ----------UPDATE Section - START-------------
-- name: UpdateUser :exec
UPDATE users
SET email = $1, username = $2, password_hash = $3
WHERE id = $4;
-- ----------UPDATE Section - END-------------//


-- ----------DELETE Section - START-------------
-- name: DropUsers :exec
DELETE FROM users;

-- ----------DELETE Section - END-------------//




