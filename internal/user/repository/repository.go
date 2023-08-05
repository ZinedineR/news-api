package repository

import (
	"context"
	"errors"
	"strings"

	"news-api/internal/user/domain"
	baseModel "news-api/pkg/db"
	"news-api/pkg/errs"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repo struct {
	db   *gorm.DB
	base *baseModel.PostgreSQLClientRepository
}

func (r repo) CreateUser(ctx context.Context, model *domain.User) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error) {
	var (
		model *domain.Verification
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Verification{}).
		Where("user_id = ?", Id).
		First(&model).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.Verified, nil
		}
		return nil, errs.Wrap(err)
	}
	return &model.Verified, nil

}

func (r repo) GetUserFullData(ctx context.Context, username, email string) (*domain.User, errs.Error) {
	var (
		models *domain.User
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.User{}).
		Where("lower(email) = ?", strings.ToLower(email)).
		Or("lower(username) = ?", strings.ToLower(username)).
		First(&models).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) StoreJWT(ctx context.Context, jwt string, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.Verification{}).
		Where("user_id = ?", Id).
		Update("jwt", jwt).Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) CheckJWT(ctx context.Context, Id uuid.UUID) (*domain.Verification, errs.Error) {
	var (
		models *domain.Verification
	)
	if err := r.db.WithContext(ctx).
		Model(&domain.Verification{}).
		Where("user_id = ?", Id).
		First(&models).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models, nil
		}
		return nil, errs.Wrap(err)
	}
	return models, nil

}

func (r repo) CreateVerification(ctx context.Context, model *domain.Verification) errs.Error {
	if err := r.db.WithContext(ctx).
		Create(&model).
		Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func (r repo) UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := r.db.WithContext(ctx).
		Model(&domain.Verification{}).
		Where("user_id = ?", Id).
		Update("verified", true).Error; err != nil {
		return errs.Wrap(err)
	}
	return nil

}

func NewRepository(db *gorm.DB, base *baseModel.PostgreSQLClientRepository) Repository {
	return &repo{db: db, base: base}
}
