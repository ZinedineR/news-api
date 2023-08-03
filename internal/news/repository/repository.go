package repository

import (
	baseModel "news-api/pkg/db"

	"gorm.io/gorm"
)

type repo struct {
	db   *gorm.DB
	base *baseModel.PostgreSQLClientRepository
}

func NewRepository(db *gorm.DB, base *baseModel.PostgreSQLClientRepository) Repository {
	return &repo{db: db, base: base}
}
