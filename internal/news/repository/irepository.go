package repository

import (
	"context"
	"news-api/internal/news/domain"
	"news-api/pkg/errs"

	"github.com/google/uuid"
)

type Repository interface {
	CreateCategories(ctx context.Context, model *domain.Categories) errs.Error
	GetDetailCategories(ctx context.Context, Id uuid.UUID) (*domain.Categories, errs.Error)
	GetCategories(ctx context.Context) (*[]domain.Categories, errs.Error)
	UpdateCategories(ctx context.Context, model *domain.Categories) errs.Error
	DeleteCategories(ctx context.Context, Id uuid.UUID) errs.Error
	SearchCategories(ctx context.Context, title string) (*domain.Categories, errs.Error)
	CreateNews(ctx context.Context, model *domain.News) errs.Error
	GetDetailNews(ctx context.Context, Id uuid.UUID) (*domain.News, errs.Error)
	GetNews(ctx context.Context) (*[]domain.News, errs.Error)
	UpdateNews(ctx context.Context, model *domain.News) errs.Error
	DeleteNews(ctx context.Context, Id uuid.UUID) errs.Error
}
