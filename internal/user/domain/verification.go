package domain

import (
	"time"

	"github.com/google/uuid"
)

const (
	VerificationTableName = "verification"
)

// VerificationModel is a model for entity.Verification
type Verification struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserId    uuid.UUID `gorm:"type:uuid;not_null" json:"user_id"`
	Expiresat time.Time `gorm:"type:date;not_null" json:"expiresat"`
	Verified  bool      `gorm:"default:false;null" json:"verified"`
	JWT       string    `gorm:"type:varchar;not_null" json:"jwt"`
	User      *User     `gorm:"foreignKey:UserId"`
}

func (model *Verification) TableName() string {
	return VerificationTableName
}
