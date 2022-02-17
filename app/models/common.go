package models

import (
	"gorm.io/gorm"
)

// 模型公用字段

// 自增ID主键
type ID struct {
	ID int64 `json:"id" gorm:"primaryKey"`
}

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
