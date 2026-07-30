// internal/base/service.go
package base

import "context"

type IBaseService[T any] interface {
	Save(ctx context.Context, entity *T) (*T, error)
	GetById(ctx context.Context, id string) (*T, error)
	Delete(ctx context.Context, id string) (bool, error)
}

type BaseService[T any] struct {
	repo BaseRepository[T]
}

func NewBaseService[T any](repo BaseRepository[T]) *BaseService[T] {
	return &BaseService[T]{repo: repo}
}

func (s *BaseService[T]) Save(ctx context.Context, entity *T) (*T, error) {
	return s.repo.Save(ctx, entity)
}

func (s *BaseService[T]) GetById(ctx context.Context, id string) (*T, error) {
	return s.repo.GetById(ctx, id)
}

func (s *BaseService[T]) Delete(ctx context.Context, id string) (bool, error) {
	return s.repo.Delete(ctx, id)
}
