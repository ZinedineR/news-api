package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	NewsTableName = "news"
)

type News struct {
	Id           uuid.UUID   `gorm:"type:uuid;primary_key;not_null" json:"id"`
	CategoriesId uuid.UUID   `gorm:"type:uuid;not_null" json:"categories_id"`
	Title        string      `gorm:"type:varchar" json:"title"`
	Description  string      `gorm:"type:varchar" json:"description"`
	Content      string      `gorm:"type:varchar" json:"content"`
	CreatedAt    time.Time   `gorm:"type:timestamp;not_null" json:"created_at"`
	UpdatedAt    time.Time   `gorm:"type:timestamp;not_null" json:"updated_at"`
	Deleted      bool        `gorm:"default:false;not_null" json:"deleted"`
	Categories   *Categories `gorm:"foreignKey:CategoriesId"`
}

type NewsDetail struct {
	Id           uuid.UUID  `gorm:"type:uuid;primary_key;not_null" json:"id"`
	CategoriesId uuid.UUID  `gorm:"type:uuid;not_null" json:"categories_id"`
	Title        string     `gorm:"type:varchar" json:"title"`
	Description  string     `gorm:"type:varchar" json:"description"`
	Content      string     `gorm:"type:varchar" json:"content"`
	CreatedAt    time.Time  `gorm:"type:timestamp;not_null" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:timestamp;not_null" json:"updated_at"`
	Deleted      bool       `gorm:"default:false;not_null" json:"deleted"`
	Comment      []*Comment `gorm:"foreignKey:PageId" json:"comment"`
}

func (model *News) CheckData() string {
	if model.Title == "" {
		return "title can't be null"
	}
	if model.Description == "" {
		return "description can't be null"
	}
	if model.Content == "" {
		return "content can't be null"
	}
	return ""
}

func (model *News) TableName() string {
	return NewsTableName
}

// func (model *News) CheckData() string {
// 	if model.Name == "" {
// 		return "Name can't be null"
// 	}
// 	if model.Email == "" {
// 		return "Email can't be null"
// 	}
// 	if model.Password == "" {
// 		return "Password can't be null"
// 	}

// 	// Check for password complexity
// 	uppercase := regexp.MustCompile(`[A-Z]`)
// 	lowercase := regexp.MustCompile(`[a-z]`)
// 	number := regexp.MustCompile(`[0-9]`)
// 	symbol := regexp.MustCompile(`[!@#$%^&*]`)

// 	if !uppercase.MatchString(model.Password) {
// 		return "Password must contain at least one uppercase letter"
// 	}
// 	if !lowercase.MatchString(model.Password) {
// 		return "Password must contain at least one lowercase letter"
// 	}
// 	if !number.MatchString(model.Password) {
// 		return "Password must contain at least one number"
// 	}
// 	if !symbol.MatchString(model.Password) {
// 		return "Password must contain at least one symbol"
// 	}
// 	if len(model.Password) < 8 {
// 		return "Password must be at least 8 characters long"
// 	}

// 	return ""
// }

// func (model *News) HashPassword(password string) error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		return err
// 	}
// 	model.Password = string(bytes)
// 	return nil
// }
// func (model *News) CheckPassword(providedPassword string) errs.Error {
// 	err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(providedPassword))
// 	if err != nil {
// 		return errs.Wrap(err)
// 	}
// 	return nil
// }
