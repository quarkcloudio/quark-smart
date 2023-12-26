package config

import "time"

type AdminConfig struct {
	Title        string                   // 应用标题
	Logo         interface{}              // Logo
	SubTitle     string                   // 登录页子标题
	IconfontUrl  string                   // 使用 IconFont 的图标配置
	Layout       string                   // layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	SplitMenus   bool                     // layout 的菜单模式为mix时，是否自动分割菜单
	ContentWidth string                   // layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	PrimaryColor string                   // 主题色,"#1890ff"
	FixedHeader  bool                     // 是否固定 header 到顶部
	FixSiderbar  bool                     // 是否固定导航
	Locale       string                   // 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	SiderWidth   int                      // 侧边菜单宽度
	Copyright    string                   // 网站版权 time.Now().Format("2006") + " QuarkGo"
	Links        []map[string]interface{} // 友情链接
}

// 后台Layout配置
var Admin = &AdminConfig{

	// 应用标题
	Title: App.Name,

	// Logo
	Logo: false,

	// 登录页子标题
	SubTitle: "信息丰富的世界里，唯一稀缺的就是人类的注意力",

	// 使用 IconFont 的图标配置
	IconfontUrl: "//at.alicdn.com/t/font_1615691_3pgkh5uyob.js",

	// layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	Layout: "mix",

	// layout 的菜单模式为mix时，是否自动分割菜单
	SplitMenus: false,

	// layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	ContentWidth: "Fluid",

	// 主题色,"#1890ff"
	PrimaryColor: "#1890ff",

	// 是否固定 header 到顶部
	FixedHeader: true,

	// 是否固定导航
	FixSiderbar: true,

	// 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	Locale: "zh-CN",

	// 侧边菜单宽度
	SiderWidth: 208,

	// 网站版权 time.Now().Format("2006") + " QuarkGo"
	Copyright: time.Now().Format("2006") + " " + App.Name,

	// 友情链接
	Links: []map[string]interface{}{
		{
			"key":   "Quark",
			"title": "Quark",
			"href":  "http://www.quarkcms.com/",
		},
		{
			"key":   "爱小圈",
			"title": "爱小圈",
			"href":  "http://www.ixiaoquan.com",
		},
		{
			"key":   "Github",
			"title": "Github",
			"href":  "https://github.com/quarkcloudio",
		},
	},
}
