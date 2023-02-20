package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 分类模型
type BannerCategory struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	Title     string         `json:"title" gorm:"size:200;not null"`
	Name      string         `json:"name" gorm:"size:100;default:null"`
	Width     int            `json:"width" gorm:"size:11;default:0;"`
	Height    int            `json:"height" gorm:"size:11;default:0;"`
	Status    int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *BannerCategory) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(21) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 21, Name: "广告管理", GuardName: "admin", Icon: "icon-banner", Type: "default", Pid: 0, Sort: 0, Path: "/banner", Show: 1, Status: 1},
		{Id: 22, Name: "广告位列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 21, Sort: 0, Path: "/api/admin/bannerCategory/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []Category{
		{Title: "首页广告位", Name: "indexPage", Status: 1},
	}
	db.Client.Create(&seeders)
}

// 获取列表
func (model *BannerCategory) List() (list []map[string]interface{}, Error error) {
	getList := []BannerCategory{}
	err := db.Client.Find(&getList).Error
	if err != nil {
		return list, err
	}

	for _, v := range getList {
		option := map[string]interface{}{
			"label": v.Title,
			"value": v.Id,
		}
		list = append(list, option)
	}

	return list, nil
}
