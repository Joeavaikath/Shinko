// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one

INSERT INTO users (id, created_at, updated_at, username, email, password_hash)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING id, username, email, password_hash, created_at, updated_at
`

type CreateUserParams struct {
	Username     string
	Email        string
	PasswordHash string
}

// ----------CREATE Section - START-------------
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Email, arg.PasswordHash)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const dropUsers = `-- name: DropUsers :exec


DELETE FROM users
`

// ----------UPDATE Section - END-------------//
// ----------DELETE Section - START-------------
func (q *Queries) DropUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, dropUsers)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one


SELECT id, username, email, password_hash, created_at, updated_at FROM users
WHERE
email = $1
`

// ----------CREATE Section - END-------------//
// ----------RETRIEVE Section - START-------------
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec


UPDATE users
SET email = $1, username = $2, password_hash = $3
WHERE id = $4
`

type UpdateUserParams struct {
	Email        string
	Username     string
	PasswordHash string
	ID           uuid.UUID
}

// ----------RETRIEVE Section - END-------------//
// ----------UPDATE Section - START-------------
func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Email,
		arg.Username,
		arg.PasswordHash,
		arg.ID,
	)
	return err
}
