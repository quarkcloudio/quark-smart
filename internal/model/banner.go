package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

// 分类模型
type Banner struct {
	Id         int            `json:"id" gorm:"autoIncrement"`
	CategoryId int            `json:"category_id" gorm:"size:11;default:0;"`
	Title      string         `json:"title" gorm:"size:200;not null"`
	UrlType    int            `json:"url_type" gorm:"size:1;not null;default:1"`
	Url        string         `json:"url" gorm:"size:200;not null"`
	Status     int            `json:"status" gorm:"size:1;not null;default:1"`
	CoverId    string         `json:"cover_id" gorm:"size:1000;default:null"`
	Sort       int            `json:"sort" gorm:"size:11;default:0;"`
	Deadline   time.Time      `json:"deadline"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Banner) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(25) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 25, Name: "广告列表", GuardName: "admin", Icon: "", Type: 2, Pid: 23, Sort: 0, Path: "/api/admin/banner/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}
