package service

import (
	"context"

	"news-api/internal/user/domain"
	"news-api/pkg/errs"

	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, model *domain.User) errs.Error
	Login(ctx context.Context, username, email string) (*domain.User, errs.Error)
	StoreJWT(ctx context.Context, jwt string, Id uuid.UUID) errs.Error
	CheckJWT(ctx context.Context, Id uuid.UUID) (*domain.Verification, errs.Error)
	CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error)
	CreateVerification(ctx context.Context, model *domain.Verification) errs.Error
	UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error
}
