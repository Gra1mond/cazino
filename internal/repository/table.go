package repository

import (
	"context"
	"errors"

	"github.com/graimond/cazino/internal/model"
	"github.com/jackc/pgx/v5"
)

type TableRepository struct {
	db *pgx.Conn
}

func NewTableReposity(db *pgx.Conn) *TableRepository {
	return &TableRepository{
		db: db,
	}
}

func (r *TableRepository) Create(ctx context.Context, name string) (model.Table, error) {
	var table model.Table
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO TABLE tables(
		name)
		VALUES $1
		RETURNING id,name`,
		name,
	).Scan(&table.ID, &table.Name)

	if err != nil {
		return model.Table{}, err
	}

	return table, nil
}

func (r *TableRepository) GetByID(ctx context.Context, id int64) (model.Table, error) {
	var table model.Table

	err := r.db.QueryRow(
		ctx,
		`SELECT id,name 
		FROM tables
		where id = $1`,
		id,
	).Scan(
		&table.ID,
		&table.Name,
	)

	if err != nil {
		return model.Table{}, err
	}
	return table, nil
}

func (r *TableRepository) List(ctx context.Context) ([]model.Table, error) {
	var result []model.Table

	rows, err := r.db.Query(
		ctx,
		`SELECT id,name
		FROM tables`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var table model.Table
		if err := rows.Scan(
			&table.ID,
			&table.Name,
		); err != nil {
			return nil, err
		}
		result = append(result, table)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, err
}

func (r *TableRepository) UpdateName(ctx context.Context, id int64, name string) (model.Table, error) {
	var table model.Table
	err := r.db.QueryRow(
		ctx,
		`UPDATE tables
		SET name = $1
		WHERE id = $2
		RETURNING id,name`,
		name,
		id,
	).Scan(&table.ID, &table.Name)

	if err != nil {
		return model.Table{}, err
	}

	return table, nil
}

func (r *TableRepository) Delete(ctx context.Context, id int64) error {
	result, err := r.db.Exec(
		ctx,
		`DELETE FROM tables
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("table not found")
	}
	return nil
}
