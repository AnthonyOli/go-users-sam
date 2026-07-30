package base

import (
	"context"
	"database/sql"
)

// INTERFACE — contrato, sem detalhe de SQL
type BaseRepository[T any] interface {
	Save(ctx context.Context, entity *T) (*T, error)
	GetById(ctx context.Context, id string) (*T, error)
	GetByEmail(ctx context.Context, email string) (*T, error)
	Delete(ctx context.Context, id string) (bool, error)
}

// STRUCT — implementação concreta, com campo DB de verdade
type SQLBaseRepository[T any] struct {
	DB    *sql.DB
	Table string
}

// método AUXILIAR, privado ao pacote (minúsculo), usado pelos métodos públicos abaixo
func (r *SQLBaseRepository[T]) exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := r.DB.ExecContext(ctx, query, args...)
	return err
}

func (r *SQLBaseRepository[T]) Delete(ctx context.Context, id string) (bool, error) {
	err := r.exec(ctx, "DELETE FROM "+r.Table+" WHERE id=$1", id) // usa o helper
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *SQLBaseRepository[T]) Save(ctx context.Context, entity *T) (*T, error) {
	// implementação real de insert/update aqui
	return entity, nil
}

func (r *SQLBaseRepository[T]) GetById(ctx context.Context, id string) (*T, error) {
	var entity T
	err := r.DB.QueryRowContext(ctx, "SELECT * FROM "+r.Table+" WHERE id=$1", id).Scan(&entity)
	return &entity, err
}
