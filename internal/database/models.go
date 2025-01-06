// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Action struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	Name             string
	Description      sql.NullString
	RecurranceWindow sql.NullInt32
	IsCalibrating    sql.NullBool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ActionEvent struct {
	ID         uuid.UUID
	ActionID   uuid.UUID
	ExecutedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Comment    sql.NullString
}

type RefreshToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
	RevokedAt sql.NullTime
}

type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
