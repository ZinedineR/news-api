package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	CommentTableName = "comment"
)

type Comment struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key;not_null" json:"id"`
	PageId    uuid.UUID `gorm:"type:uuid;not_null" json:"page_id"`
	Name      string    `gorm:"type:varchar" json:"name"`
	Comment   string    `gorm:"type:varchar" json:"comment"`
	CreatedAt time.Time `gorm:"type:timestamp;not_null" json:"created_at"`
	News      *News     `gorm:"foreignKey:PageId"`
}

func (model *Comment) CheckData() string {
	if model.Comment == "" {
		return "comment can't be null"
	}
	return ""
}

func (model *Comment) TableName() string {
	return CommentTableName
}

// func (model *Comment) CheckData() string {
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

// func (model *Comment) HashPassword(password string) error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		return err
// 	}
// 	model.Password = string(bytes)
// 	return nil
// }
// func (model *Comment) CheckPassword(providedPassword string) errs.Error {
// 	err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(providedPassword))
// 	if err != nil {
// 		return errs.Wrap(err)
// 	}
// 	return nil
// }
