package store

import "gorm.io/plugin/soft_delete"

type Base struct {
	ID        int                   `json:"id" gorm:"primaryKey;type:serial;"`
	CreatedAt int                   `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int                   `json:"updated_at" gorm:"autoUpdateTime:milli"`
	Deleted   soft_delete.DeletedAt `json:"deleted" gorm:"softDelete:flag"`
}
