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

func (s service) CreateNews(ctx context.Context, model *domain.News) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.CreatedAt = time.Now()
		model.UpdatedAt = time.Now()
	}

	if err := s.newsRepo.CreateNews(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) GetDetailNews(ctx context.Context, Id uuid.UUID) (*domain.NewsDetail, errs.Error) {
	result, err := s.newsRepo.GetDetailNews(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) GetNews(ctx context.Context) (*[]domain.News, errs.Error) {
	result, err := s.newsRepo.GetNews(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) UpdateNews(ctx context.Context, model *domain.News) errs.Error {
	model.UpdatedAt = time.Now()
	if err := s.newsRepo.UpdateNews(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) DeleteNews(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := s.newsRepo.DeleteNews(ctx, Id); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) CreateCustom(ctx context.Context, model *domain.Custom) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.CreatedAt = time.Now()
		model.UpdatedAt = time.Now()
	}

	if err := s.newsRepo.CreateCustom(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) GetDetailCustom(ctx context.Context, Id uuid.UUID) (*domain.Custom, errs.Error) {
	result, err := s.newsRepo.GetDetailCustom(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) GetCustom(ctx context.Context) (*[]domain.Custom, errs.Error) {
	result, err := s.newsRepo.GetCustom(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) UpdateCustom(ctx context.Context, model *domain.Custom) errs.Error {
	model.UpdatedAt = time.Now()
	if err := s.newsRepo.UpdateCustom(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) DeleteCustom(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := s.newsRepo.DeleteCustom(ctx, Id); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) CreateComment(ctx context.Context, model *domain.Comment) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.CreatedAt = time.Now()
	}
	if model.Name == "" {
		model.Name = "Anonymous"
	}

	if err := s.newsRepo.CreateComment(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}
