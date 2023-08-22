package service_test

import (
	"errors"
	"news-api/internal/base/app"
	"news-api/internal/user/domain"
	userMocks "news-api/internal/user/mocks"
	"news-api/internal/user/service"
	"news-api/pkg/errs"
	"testing"

	"github.com/google/uuid"
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

func TestStoreJwt(t *testing.T) {
	mockAppCtx := &app.Context{}
	mockJwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	mockUUID := uuid.New()
	t.Run("user service StoreJWT success", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("StoreJWT", mockAppCtx, mockJwt, mockUUID).Return(nil)
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.StoreJWT(mockAppCtx, mockJwt, mockUUID)
		assert.Nil(t, err)
	})

	t.Run("user service StoreJWT failed", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("StoreJWT", mockAppCtx, mockJwt, mockUUID).Return(errs.Wrap(errors.New("createuser failed, database problem")))
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.StoreJWT(mockAppCtx, mockJwt, mockUUID)
		assert.NotNil(t, err)
	})
}

func TestCheckJWT(t *testing.T) {
	mockAppCtx := &app.Context{}
	mockJwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	mockUUID := uuid.New()
	successCheck := &domain.Verification{
		Jwt:    mockJwt,
		UserId: mockUUID,
	}
	t.Run("user service CheckJWT success", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CheckJWT", mockAppCtx, mockUUID).Return(successCheck, nil)
		mockUserService := service.NewService(mockUserRepo)
		mockJWT, err := mockUserService.CheckJWT(mockAppCtx, mockUUID)
		assert.Equal(t, successCheck.Jwt, mockJWT.Jwt)
		assert.Nil(t, err)
	})

	t.Run("user service CheckJWT failed", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CheckJWT", mockAppCtx, mockUUID).Return(nil, errs.Wrap(errors.New("jwt not found")))
		mockUserService := service.NewService(mockUserRepo)
		mockJWT, err := mockUserService.CheckJWT(mockAppCtx, mockUUID)
		assert.Nil(t, mockJWT)
		assert.NotNil(t, err)
	})
}

func TestCheckVerified(t *testing.T) {
	mockAppCtx := &app.Context{}
	mockUUID := uuid.New()
	expectedResult := true
	t.Run("user service CheckVerified success", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CheckVerified", mockAppCtx, mockUUID).Return(&expectedResult, nil)
		mockUserService := service.NewService(mockUserRepo)
		result, err := mockUserService.CheckVerified(mockAppCtx, mockUUID)
		assert.Nil(t, err)
		assert.Equal(t, expectedResult, *result)
	})

	t.Run("user service CheckVerified failed", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CheckVerified", mockAppCtx, mockUUID).Return(nil, errs.Wrap(errors.New("verified not found")))
		mockUserService := service.NewService(mockUserRepo)
		mockVerified, err := mockUserService.CheckVerified(mockAppCtx, mockUUID)
		assert.Nil(t, mockVerified)
		assert.NotNil(t, err)
	})
}

func TestCreateVerification(t *testing.T) {
	mockAppCtx := &app.Context{}
	mockVerification := &domain.Verification{}
	t.Run("user service CreateVerification success", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CreateVerification", mockAppCtx, mockVerification).Return(nil)
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.CreateVerification(mockAppCtx, mockVerification)
		assert.Nil(t, err)
	})

	t.Run("user service CreateVerification failed", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("CreateVerification", mockAppCtx, mockVerification).Return(errs.Wrap(errors.New("createuser failed, database problem")))
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.CreateVerification(mockAppCtx, mockVerification)
		assert.NotNil(t, err)
	})
}

func TestUpdateVerification(t *testing.T) {
	mockAppCtx := &app.Context{}
	mockUUID := uuid.New()
	t.Run("user service UpdateVerification success", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("UpdateVerification", mockAppCtx, mockUUID).Return(nil)
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.UpdateVerification(mockAppCtx, mockUUID)
		assert.Nil(t, err)
	})

	t.Run("user service UpdateVerification failed", func(t *testing.T) {
		mockUserRepo := new(userMocks.Repository)
		mockUserRepo.On("UpdateVerification", mockAppCtx, mockUUID).Return(errs.Wrap(errors.New("Updateuser failed, database problem")))
		mockUserService := service.NewService(mockUserRepo)
		err := mockUserService.UpdateVerification(mockAppCtx, mockUUID)
		assert.NotNil(t, err)
	})
}
