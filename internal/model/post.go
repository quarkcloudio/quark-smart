package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
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
func (m *Post) Seeder() {

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 18, Name: "内容管理", GuardName: "admin", Icon: "icon-read", Type: "default", Pid: 0, Sort: 0, Path: "/post", Show: 1, Status: 1},
		{Id: 19, Name: "文章列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 18, Sort: 0, Path: "/api/admin/post/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认文章
	postSeeders := []Post{
		{Title: "Hello world!", Content: "Hello world!", Status: 1},
	}
	db.Client.Create(&postSeeders)
}
