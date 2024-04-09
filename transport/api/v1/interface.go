package v1

import (
	"context"
	userStructs "newsletter_backend_api/transport/api/v1/structs/user"

	"newsletter_backend_api/pkg/id"
	svcmodel "newsletter_backend_api/service/model"
)

type Service interface {
	CreateUser(ctx context.Context, user userStructs.CreateUserStruct) error
	ListUsers(ctx context.Context) ([]svcmodel.User, error)
	GetUser(ctx context.Context, userID id.User) (*svcmodel.User, error)
	DeleteUser(ctx context.Context, userID id.User) error
}
