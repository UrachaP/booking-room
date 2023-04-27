package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        int            `json:"id"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
}

type Pagination struct {
	Page  int    `json:"page" query:"page"`
	Limit int    `json:"limit" query:"limit"`
	Sort  string `json:"sort" query:"sort"`
}
