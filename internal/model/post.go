package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-go/pkg/lister"
	"gorm.io/gorm"
)

// 文章模型
type Post struct {
	Id            int            `json:"id" gorm:"autoIncrement"`
	Adminid       int            `json:"adminid"`
	Uid           int            `json:"uid"`
	CategoryId    int            `json:"category_id"`
	Tags          string         `json:"tags" gorm:"size:200;default:null"`
	Title         string         `json:"title" gorm:"size:200;not null"`
	Name          string         `json:"name" gorm:"size:200;default:null"`
	Author        string         `json:"author" gorm:"size:200;default:null"`
	Source        string         `json:"source" gorm:"size:200;default:null"`
	Description   string         `json:"description" gorm:"size:200;default:null"`
	Password      string         `json:"password" gorm:"size:200;default:null"`
	CoverIds      string         `json:"cover_ids" gorm:"size:1000;default:null"`
	Pid           int            `json:"pid" gorm:"default:0"`
	Level         int            `json:"level" gorm:"size:11;default:0"`
	Type          string         `json:"type" gorm:"size:200;not null;default:ARTICLE"`
	ShowType      int            `json:"show_type" gorm:"size:4;default:0"`
	Position      string         `json:"position" gorm:"size:100;default:null"`
	Link          string         `json:"link" gorm:"size:100;default:null"`
	Content       string         `json:"content" gorm:"size:5000;default:null"`
	Comment       int            `json:"comment" gorm:"default:0"`
	View          int            `json:"view" gorm:"default:0"`
	PageTpl       string         `json:"page_tpl" gorm:"size:100"`
	CommentStatus int            `json:"comment_status" gorm:"size:1;not null;default:0"`
	FileIds       string         `json:"file_ids" gorm:"size:1000;default:null"`
	Status        int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Post) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(18) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 18, Name: "内容管理", GuardName: "admin", Icon: "icon-read", Type: "default", Pid: 0, Sort: 0, Path: "/post", Show: 1, Status: 1},
		{Id: 20, Name: "文章列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 18, Sort: 0, Path: "/api/admin/article/index", Show: 1, Status: 1},
		{Id: 21, Name: "单页管理", GuardName: "admin", Icon: "icon-page", Type: "default", Pid: 0, Sort: 0, Path: "/page", Show: 1, Status: 1},
		{Id: 22, Name: "单页列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 21, Sort: 0, Path: "/api/admin/page/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认内容
	seeders := []Post{
		{Title: "关于我们", Name: "aboutus", Content: "关于我们", Status: 1, Type: "PAGE"},
	}
	db.Client.Create(&seeders)
}

// 获取菜单的有序列表
func (model *Post) OrderedList(root bool) (list []map[string]interface{}, Error error) {
	var data []map[string]interface{}
	err := db.Client.
		Model(&model).
		Where("type", "PAGE").
		Order("id asc").
		Find(&data).Error
	if err != nil {
		return list, err
	}

	trees, err := lister.ListToTree(data, "id", "pid", "children", 0)
	if err != nil {
		return list, err
	}

	treeList, err := lister.TreeToOrderedList(trees, 0, "title", "children")
	if err != nil {
		return list, err
	}

	// 是否有跟节点
	if root {
		list = append(list, map[string]interface{}{
			"label": "根节点",
			"value": 0,
		})
	}

	for _, v := range treeList {
		option := map[string]interface{}{
			"label": v.((map[string]interface{}))["title"],
			"value": v.(map[string]interface{})["id"],
		}
		list = append(list, option)
	}

	return list, nil
}
