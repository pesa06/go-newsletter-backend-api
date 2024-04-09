package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"newsletter_backend_api/pkg/id"
	"newsletter_backend_api/service/model"
	"newsletter_backend_api/transport/api/v1/structs/user"
)

func (s Service) ListUsers(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetUser(ctx context.Context, userID id.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteUser(ctx context.Context, userID id.User) error {
	//TODO implement me
	panic("implement me")
}

// create user function
func (s Service) CreateUser(ctx context.Context, userRequest user.CreateUserStruct) error {
	userModel := prepareUserModelFromRequest(userRequest)
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	slog.Info("hashedPw", string(hashedPw))
	userModel.Password = string(hashedPw)
	return s.repository.CreateUser(ctx, userModel)
}

func prepareUserModelFromRequest(userRequest user.CreateUserStruct) model.User {
	return model.User{
		Username: userRequest.Username,
		Role:     userRequest.Role,
	}
}
