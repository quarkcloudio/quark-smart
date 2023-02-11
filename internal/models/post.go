package models

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 字段
type Post struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	Title     string         `json:"title" gorm:"size:200;not null"`
	Content   string         `json:"content" gorm:"size:5000"`
	Status    int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (model *Post) Seeder() {
	seeders := []Post{
		{Title: "Hello world!", Content: "Hello world!", Status: 1},
	}

	db.Client.Create(&seeders)
}
