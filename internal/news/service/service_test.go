package service_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"news-api/internal/base/app"
	"news-api/internal/news/domain"
	userMocks "news-api/internal/news/mocks"
	"news-api/internal/news/service"
	"news-api/pkg/errs"
	"testing"
)

func TestNewService(t *testing.T) {
	t.Run("test create NewService service", func(t *testing.T) {
		mockNewsRepo := new(userMocks.Repository)
		mockNewsService := service.NewService(mockNewsRepo)
		assert.NotNil(t, mockNewsService)
	})
}

func TestCreateCategoriesCreateCategories(t *testing.T) {
	mockAppCtx := &app.Context{}
	mockCategories := &domain.Categories{}
	t.Run("user service CreateCategories success", func(t *testing.T) {
		mockNewsRepo := new(userMocks.Repository)
		mockNewsRepo.On("CreateCategories", mockAppCtx, mockCategories).Return(nil)
		mockNewsService := service.NewService(mockNewsRepo)
		err := mockNewsService.CreateCategories(mockAppCtx, mockCategories)
		assert.Nil(t, err)
	})

	t.Run("user service CreateCategories failed", func(t *testing.T) {
		mockNewsRepo := new(userMocks.Repository)
		mockNewsRepo.On("CreateCategories", mockAppCtx, mockCategories).Return(errs.Wrap(errors.New("createuser failed, database problem")))
		mockNewsService := service.NewService(mockNewsRepo)
		err := mockNewsService.CreateCategories(mockAppCtx, mockCategories)
		assert.NotNil(t, err)
	})
}
