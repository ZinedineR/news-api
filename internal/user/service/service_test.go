package service_test

import (
	"errors"
	"news-api/internal/base/app"
	"news-api/internal/user/domain"
	userMocks "news-api/internal/user/mocks"
	"news-api/internal/user/service"
	"news-api/pkg/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

// var mockRepository = mocks.Repository{Mock: mock.Mock{}}
// var mockService = service{profileRepo: &mockRepository}
// var ctx = context.Background()

func TestNewService(t *testing.T) {
	t.Run("test create NewService service", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserService := service.NewService(mockUserRepo)
		assert.NotNil(t, mockUserService)
	})
}

func TestCreateUser(t *testing.T) {
	mockAppCtx := &app.Context{}
	mockUser := &domain.User{}
	t.Run("user service createuser success", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CreateUser", mockAppCtx, mockUser).Return(nil)
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.CreateUser(mockAppCtx, mockUser)
		assert.Nil(t, err)
	})

	t.Run("user service createuser failed", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CreateUser", mockAppCtx, mockUser).Return(errs.Wrap(errors.New("createuser failed, database problem")))
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.CreateUser(mockAppCtx, mockUser)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	mockAppCtx := &app.Context{}
	successLogin := &domain.User{
		Username: "success_username",
		Email:    "success@email.com",
	}
	t.Run("user service login success", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("GetUserFullData", mockAppCtx, "success_username", "success@email.com").Return(successLogin, nil)
		mockUserService := service.NewService(mockUserRepo)
		mockUser, err := mockUserService.Login(mockAppCtx, "success_username", "success@email.com")
		assert.Equal(t, successLogin.Username, mockUser.Username)
		assert.Nil(t, err)
	})

	t.Run("user service login failed", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("GetUserFullData", mockAppCtx, "unmatched_username", "unmatched@email.com").Return(nil, errs.Wrap(errors.New("login not found")))
		mockUserService := service.NewService(mockUserRepo)
		mockUser, err := mockUserService.Login(mockAppCtx, "unmatched_username", "unmatched@email.com")
		assert.Nil(t, mockUser)
		assert.NotNil(t, err)
	})
}
