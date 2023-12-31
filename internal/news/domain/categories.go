package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	CategoriesTableName = "categories"
)

type Categories struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;not_null" json:"id"`
	Title     string    `gorm:"type:varchar" json:"title"`
	CreatedAt time.Time `gorm:"type:date;not_null" json:"created_at"`
	Deleted   bool      `gorm:"default:false;not_null" json:"deleted"`
}

func (model *Categories) CheckData() string {
	if model.Title == "" {
		return "title can't be null"
	}
	return ""
}

func (model *Categories) TableName() string {
	return CategoriesTableName
}
