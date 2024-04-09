package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"newsletter_backend_api/pkg/id"
	"newsletter_backend_api/repository/sql/query"
	"newsletter_backend_api/service/model"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}

}

func (r *UserRepository) ReadUser(ctx context.Context, userID id.User) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) ListUser(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) error {
	//store user into database
	_, err := r.pool.Query(ctx, query.CreateUser, pgx.NamedArgs{
		"username": user.Username,
		"password": user.Password,
		"role":     user.Role,
	})
	if err != nil {
		return err
	}
	return nil

}
