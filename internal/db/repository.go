package db

import (
	"context"
	"go-serverless/internal/db/entities"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (*entities.User, error)
}
