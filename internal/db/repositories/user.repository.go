package repositories

import (
	"context"
	"go-serverless/helpers"
	"go-serverless/internal/base"
	"go-serverless/internal/db/entities"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	base.SQLBaseRepository[entities.User]
}

func (r *UserRepository) Save(ctx context.Context, user *entities.User) (*entities.User, error) {
	var row pgx.Row

	if user.Id == "" {
		hashedPassword, hashError := helpers.HashPassword(user.PasswordHash)
		if hashError != nil {
			return nil, hashError
		}
		user.PasswordHash = hashedPassword

		row = r.Pool.QueryRow(ctx,
			`INSERT INTO users (name, email, password_hash, phone, is_active)
			 VALUES ($1, $2, $3, $4, $5)
			 RETURNING id, name, email, phone, is_active, created_at, updated_at`,
			user.Name, user.Email, user.PasswordHash, user.Phone, user.IsActive)
	} else {
		row = r.Pool.QueryRow(ctx,
			`UPDATE users SET name=$1, email=$2, phone=$3, is_active=$4
			 WHERE id=$5
			 RETURNING id, name, email, phone, is_active, created_at, updated_at`,
			user.Name, user.Email, user.Phone, user.IsActive, user.Id)
	}

	var u entities.User
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.Phone, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepository) UpdateUserPassword(ctx context.Context, userId string, newPassword string) error {
	//TODO : Implement password update logic. Should be used in password reset flow and change password flow.
	return nil
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
	err := row.Scan(&u.Id, &u.Name, &u.Email, &u.PasswordHash, &u.Phone, &u.IsActive, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
