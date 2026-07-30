package repositories

import (
	"context"
	"database/sql"
	"go-serverless/internal/base"
	"go-serverless/internal/db/entities"
)

type UserRepository struct {
	base.SQLBaseRepository[entities.User]
}

func NewUserRepository(sqlDb *sql.DB) *UserRepository {
	return &UserRepository{
		SQLBaseRepository: base.SQLBaseRepository[entities.User]{
			DB:    sqlDb,
			Table: "users",
		},
	}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	row := r.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE email=$1", email)
	var u entities.User
	if err := row.Scan(&u.ID, &u.Email); err != nil {
		return nil, err
	}
	return &u, nil
}
