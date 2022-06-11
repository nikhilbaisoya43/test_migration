package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Base struct {
	ID        string     `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedAt", time.Now().Format(time.RFC3339))
	scope.SetColumn("UpdatedAt", time.Now().Format(time.RFC3339))
}
