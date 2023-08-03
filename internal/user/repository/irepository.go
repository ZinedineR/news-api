package repository

import (
	"context"

	"news-api/internal/user/domain"
	"news-api/pkg/errs"

	"github.com/google/uuid"
)

type Repository interface {
	StoreJWT(ctx context.Context, jwt string, Id uuid.UUID) errs.Error
	CheckJWT(ctx context.Context, Id uuid.UUID) (*domain.User, errs.Error)
	CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error)
	GetUserFullData(ctx context.Context, email string) (*domain.User, errs.Error)
	CreateUser(ctx context.Context, model *domain.User) errs.Error
	CreateVerification(ctx context.Context, model *domain.Verification) errs.Error
	UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error
}
