package store

import "gorm.io/gorm"

type Base struct {
	ID        int            `json:"id" gorm:"primaryKey;type:serial;"`
	CreatedAt int            `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int            `json:"updated_at" gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
