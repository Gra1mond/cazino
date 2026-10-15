package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct{
	db *pgx.Conn
}

func NewUserRepositor(db *pgx.Conn)(*UserRepository){
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository)Create(ctx context.Context,user )