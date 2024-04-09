package service

import (
	"context"

	"newsletter_backend_api/pkg/id"
	"newsletter_backend_api/service/model"
)

type Repository interface {
	ReadUser(ctx context.Context, userID id.User) (*model.User, error)
	ListUser(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user model.User) error
}

type Service struct {
	repository Repository
}

func NewService(
	repository Repository,
) (Service, error) {
	return Service{
		repository: repository,
	}, nil
}
