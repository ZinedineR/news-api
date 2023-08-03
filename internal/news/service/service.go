package service

import (
	"news-api/internal/user/repository"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{authRepo: repo}
}

type service struct {
	authRepo repository.Repository
}
