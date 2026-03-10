package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, email, passwordHash string) error
}

type AuthService struct {
	repo UserRepo
}

func NewAuthService(repo UserRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, email, password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.CreateUser(ctx, email, string(hash))
}
