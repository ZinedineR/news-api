package service

import (
	"context"
	"news-api/internal/news/domain"
	"news-api/internal/news/repository"
	"news-api/pkg/errs"
	"time"

	"github.com/google/uuid"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{newsRepo: repo}
}

type service struct {
	newsRepo repository.Repository
}

func (s service) CreateCategories(ctx context.Context, model *domain.Categories) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.CreatedAt = time.Now()
	}

	if err := s.newsRepo.CreateCategories(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) GetDetailCategories(ctx context.Context, Id uuid.UUID) (*domain.Categories, errs.Error) {
	result, err := s.newsRepo.GetDetailCategories(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) GetCategories(ctx context.Context) (*[]domain.Categories, errs.Error) {
	result, err := s.newsRepo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) UpdateCategories(ctx context.Context, model *domain.Categories) errs.Error {
	if err := s.newsRepo.UpdateCategories(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) DeleteCategories(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := s.newsRepo.DeleteCategories(ctx, Id); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) SearchCategories(ctx context.Context, title string) (*domain.Categories, errs.Error) {
	result, err := s.newsRepo.SearchCategories(ctx, title)
	if err != nil {
		return nil, err
	}
	return result, nil
}
