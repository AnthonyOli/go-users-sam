package base

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	Pool  *pgxpool.Pool
	Table string
}

type PaginatedResponse[T any] struct {
	Total    int  `json:"total"`
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
	Data     []*T `json:"data"`
}

func (r *SQLBaseRepository[T]) List(ctx context.Context, pageSize *int, page *int) (*PaginatedResponse[T], error) {
	var response PaginatedResponse[T]

	response.PageSize = 10

	if pageSize != nil {
		response.PageSize = *pageSize
	}

	response.Page = 1
	if page != nil {
		response.Page = *page
	}
	offset := (response.Page - 1) * response.PageSize

	rows, err := r.Pool.Query(ctx, "SELECT * FROM "+r.Table+" LIMIT $1 OFFSET $2 ", response.PageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entities, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	count, err := r.Count(ctx)
	if err != nil {
		return nil, err
	}

	response.Total = count
	response.Data = entities

	return &response, nil
}

func (r *SQLBaseRepository[T]) Count(ctx context.Context) (int, error) {
	var total int

	err := r.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM "+r.Table).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

// método AUXILIAR, privado ao pacote (minúsculo), usado pelos métodos públicos abaixo
func (r *SQLBaseRepository[T]) exec(ctx context.Context, query string, args ...interface{}) error {
	_, err := r.Pool.Exec(ctx, query, args...)
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
	rows, err := r.Pool.Query(ctx, "SELECT * FROM "+r.Table+" WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entity, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[T])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return entity, nil
}
