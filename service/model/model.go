package model

import (
	"newsletter_backend_api/pkg/id"
)

type User struct {
	ID       id.User
	Username string
	Password string
	Role     string
}

const (
	Publisher  = "publisher"
	Subscriber = "subscriber"
)
