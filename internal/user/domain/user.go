package domain

import (
	"news-api/pkg/errs"
	"regexp"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserTableName = "user"
)

type User struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name      string    `gorm:"type:varchar;not_null" json:"name"`
	Email     string    `gorm:"type:varchar;not_null" json:"email"`
	Password  string    `gorm:"type:varchar;not_null" json:"password"`
	CreatedAt time.Time `gorm:"type:date;not_null" json:"created_at"`
}

type UserData struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name      string    `gorm:"type:varchar;not_null" json:"name"`
	CreatedAt time.Time `gorm:"type:date;not_null" json:"created_at"`
}

type UserLogin struct {
	Email    string `gorm:"type:varchar;not_null" json:"email"`
	Password string `gorm:"type:varchar;not_null" json:"password"`
}

func (model *User) TableName() string {
	return UserTableName
}

func (model *User) CheckData() string {
	if model.Name == "" {
		return "Name can't be null"
	}
	if model.Email == "" {
		return "Email can't be null"
	}
	if model.Password == "" {
		return "Password can't be null"
	}

	// Check for password complexity
	uppercase := regexp.MustCompile(`[A-Z]`)
	lowercase := regexp.MustCompile(`[a-z]`)
	number := regexp.MustCompile(`[0-9]`)
	symbol := regexp.MustCompile(`[!@#$%^&*]`)

	if !uppercase.MatchString(model.Password) {
		return "Password must contain at least one uppercase letter"
	}
	if !lowercase.MatchString(model.Password) {
		return "Password must contain at least one lowercase letter"
	}
	if !number.MatchString(model.Password) {
		return "Password must contain at least one number"
	}
	if !symbol.MatchString(model.Password) {
		return "Password must contain at least one symbol"
	}
	if len(model.Password) < 8 {
		return "Password must be at least 8 characters long"
	}

	return ""
}

func (model *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	model.Password = string(bytes)
	return nil
}
func (model *User) CheckPassword(providedPassword string) errs.Error {
	err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(providedPassword))
	if err != nil {
		return errs.Wrap(err)
	}
	return nil
}
