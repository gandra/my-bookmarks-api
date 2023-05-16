package domain

import (
	"github.com/gandra/my-bookmarks/utils/date_utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// https://gorm.io/docs/hooks.html

// Base contains common columns for all tables.
type Base struct {
	Id        uuid.UUID `json:"id", gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate will set a UUID and CreatedAt
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.Id = uuid.New()
	base.CreatedAt = date_utils.GetNow()
	return nil
}

// BeforeUpdate will update UpdatedAt
func (base *Base) BeforeUpdate(tx *gorm.DB) error {
	base.UpdatedAt = date_utils.GetNow()
	return nil
}
