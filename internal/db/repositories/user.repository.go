package repositories

import (
	"context"
	"go-serverless/internal/base"
	"go-serverless/internal/db/entities"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	base.SQLBaseRepository[entities.User]
}

func NewUserRepository(dbPool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		SQLBaseRepository: base.SQLBaseRepository[entities.User]{
			Pool:  dbPool,
			Table: "users",
		},
	}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	row := r.Pool.QueryRow(ctx, "SELECT * FROM users WHERE email=$1", email)
	var u entities.User
	if err := row.Scan(&u.Id, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}
