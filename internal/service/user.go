package service

import (
	"context"
	"errors"

	"github.com/graimond/cazino/internal/auth"
	"github.com/graimond/cazino/internal/model"
	"github.com/graimond/cazino/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context,
	name string, email string, password string) (model.User, error) {
	// add name generation from slice
	if name == "" {
		return model.User{}, errors.New("name is reqired")
	}
	if email == "" {
		return model.User{}, errors.New("email is required")
	}
	if password == "" {
		return model.User{}, errors.New("password is reqired")
	}
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return model.User{}, err
	}
	user := model.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	}
	return s.repo.Create(ctx, user)
}

func (s *UserService) Login(ctx context.Context,
	email string, name string, password string) (string, error) {
	if email == "" && name == "" {
		return "", errors.New("email or name is reqiered")
	}
	if password == "" {
		return "", errors.New("password is reqired")
	}
	var (
		user model.User
		err  error
	)

	if email == "" {
		user, err = s.repo.GetByName(ctx, name)
	} else {
		user, err = s.repo.GetByEmail(ctx, email)
	}

	if err != nil {
		return "", errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return "", errors.New("ivalid email or password")
	}
	token, err := auth.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
