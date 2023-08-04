package repository

import (
	"context"
	"errors"
	"news-api/internal/news/domain"
	baseModel "news-api/pkg/db"
	"news-api/pkg/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	db   *gorm.DB
	base *baseModel.PostgreSQLClientRepository
}

func NewRepository(db *gorm.DB, base *baseModel.PostgreSQLClientRepository) Repository {
	return &repo{db: db, base: base}
}

func (r repo) CreateCategories(ctx context.Context, model *domain.Categories) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) GetDetailCategories(ctx context.Context, Id uuid.UUID) (*domain.Categories, errs.Error) {
	var (
		models *domain.Categories
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Categories{}).
		First(&models, Id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil
}

func (r repo) GetCategories(ctx context.Context) (*[]domain.Categories, errs.Error) {
	var (
		models *[]domain.Categories
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Categories{}).
		Where("deleted", false).
		Find(&models).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) UpdateCategories(ctx context.Context, model *domain.Categories) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.Categories{Id: model.Id}).
		Select("title").
		Updates(model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) DeleteCategories(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.Categories{Id: Id}).
		Update("deleted", true).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) SearchCategories(ctx context.Context, title string) (*domain.Categories, errs.Error) {
	var (
		models *domain.Categories
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Categories{}).
		Where("title", title).
		First(&models).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil
}

func (r repo) CreateNews(ctx context.Context, model *domain.News) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) GetDetailNews(ctx context.Context, Id uuid.UUID) (*domain.News, errs.Error) {
	var (
		models *domain.News
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.News{}).
		First(&models, Id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil
}

func (r repo) GetNews(ctx context.Context) (*[]domain.News, errs.Error) {
	var (
		models *[]domain.News
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.News{}).
		Where("deleted", false).
		Find(&models).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) UpdateNews(ctx context.Context, model *domain.News) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.News{Id: model.Id}).
		Select("categories_id, title, description, content, updated_at").
		Updates(model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) DeleteNews(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.News{Id: Id}).
		Update("deleted", true).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) CreateCustom(ctx context.Context, model *domain.Custom) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) GetDetailCustom(ctx context.Context, Id uuid.UUID) (*domain.Custom, errs.Error) {
	var (
		models *domain.Custom
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Custom{}).
		First(&models, Id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil
}

func (r repo) GetCustom(ctx context.Context) (*[]domain.Custom, errs.Error) {
	var (
		models *[]domain.Custom
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Custom{}).
		Where("deleted", false).
		Find(&models).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) UpdateCustom(ctx context.Context, model *domain.Custom) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.Custom{Id: model.Id}).
		Select("categories_id, title, description, content, updated_at").
		Updates(model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (r repo) DeleteCustom(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.Custom{Id: Id}).
		Update("deleted", true).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil
}
