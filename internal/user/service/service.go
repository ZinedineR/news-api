package service

import (
	"context"
	"time"

	"news-api/internal/user/domain"
	"news-api/internal/user/repository"
	"news-api/pkg/errs"

	"github.com/google/uuid"
)

// NewService creates new user service
func NewService(repo repository.Repository) Service {
	return &service{profileRepo: repo}
}

type service struct {
	profileRepo repository.Repository
}

func (s service) CreateUser(ctx context.Context, model *domain.User) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.CreatedAt = time.Now()
	}

	if err := s.profileRepo.CreateUser(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) Login(ctx context.Context, username, email string) (*domain.User, errs.Error) {
	result, err := s.profileRepo.GetUserFullData(ctx, username, email)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) StoreJWT(ctx context.Context, jwt string, Id uuid.UUID) errs.Error {
	if err := s.profileRepo.StoreJWT(ctx, jwt, Id); err != nil {
		return err
	}
	return nil
}

func (s service) CheckJWT(ctx context.Context, Id uuid.UUID) (*domain.Verification, errs.Error) {
	result, err := s.profileRepo.CheckJWT(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) CheckVerified(ctx context.Context, Id uuid.UUID) (*bool, errs.Error) {
	result, err := s.profileRepo.CheckVerified(ctx, Id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s service) CreateVerification(ctx context.Context, model *domain.Verification) errs.Error {
	if model.Id == uuid.Nil {
		model.Id = uuid.New()
		model.Expiresat = time.Now().Add(time.Hour * 24)
	}

	if err := s.profileRepo.CreateVerification(ctx, model); err != nil {
		return errs.Wrap(err)
	}
	return nil
}

func (s service) UpdateVerification(ctx context.Context, Id uuid.UUID) errs.Error {
	if err := s.profileRepo.UpdateVerification(ctx, Id); err != nil {
		return errs.Wrap(err)
	}
	return nil
}
