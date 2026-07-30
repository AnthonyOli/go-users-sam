package user

import (
	"context"
	"go-serverless/internal/base"
	"go-serverless/internal/db/entities"
)

type UserService struct {
	repo base.BaseRepository[entities.User]
}

func NewUserService(repo base.BaseRepository[entities.User]) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Login(ctx context.Context, email string) (*entities.User, error) {
	return s.repo.GetByEmail(ctx, email)
}
