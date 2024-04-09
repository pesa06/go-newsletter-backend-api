package model

import (
	"time"

	"newsletter_backend_api/pkg/id"
)

type User struct {
	ID        id.User   `db:"id"`
	Username  string    `db:"username"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
