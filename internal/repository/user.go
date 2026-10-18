package repository

import (
	"context"

	"github.com/graimond/cazino/internal/model"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	err := r.db.QueryRow(
		ctx,
		`INSERT INTO TABLE users(
		name,email,password_hash)
		VALUES $1,$2,$3
		RETURNING id`,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(&user.ID)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(
		ctx,
		`SELECT id,name,email,password_hash
		FROM users
		WHERE email = $1`,
		email,
	).Scan(&user)
	return user, err
}

func (r *UserRepository) GetByName(ctx context.Context, name string) (model.User, error) {
	var user model.User

	err := r.db.QueryRow(
		ctx,
		`SELECT id,name,email,password_hash
		FROM users
		WHERE name = $1`,
		name,
	).Scan(&user)
	return user, err
}
