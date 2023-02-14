package config

import "time"

const APP_LOGO = false
const ICONFONT_URL = "//at.alicdn.com/t/font_1615691_3pgkh5uyob.js"

var Admin = map[string]interface{}{
	// QuarkAdmin App Name
	"name": App["name"],

	// QuarkAdmin Version
	"version": "0.0.1",

	// QuarkAdmin logo
	"logo": APP_LOGO,

	// The description of login page.
	"description": "信息丰富的世界里，唯一稀缺的就是人类的注意力",

	// 使用 IconFont 的图标配置
	"iconfont_url": ICONFONT_URL,

	// The layout of QuarkAdmin
	"layout": map[string]interface{}{

		// layout 的左上角 的 title
		"title": App["name"],

		// layout 的左上角 的 logo
		"logo": APP_LOGO,

		// layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
		"layout": "side",

		// layout 的菜单模式为mix时，是否自动分割菜单
		"split_menus": false,

		// layout 的菜单模式为mix时，顶部主题 'dark' | 'light'
		"header_theme": "dark",

		// layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
		"content_width": "Fluid",

		// 导航的主题，'light' | 'dark'
		"nav_theme": "dark",

		// 主题色
		"primary_color": "#1890ff",

		// 是否固定 header 到顶部
		"fixed_header": true,

		// 是否固定导航
		"fix_siderbar": true,

		// 使用 IconFont 的图标配置
		"iconfont_url": ICONFONT_URL,

		// 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
		"locale": "zh-CN",

		// 侧边菜单宽度
		"sider_width": 208,
	},

	// 网站版权
	"copyright": time.Now().Format("2006") + " QuarkGo",

	// 友情链接
	"links": []map[string]interface{}{
		{
			"title": "Quark",
			"href":  "http://www.quarkcms.com/",
		},
		{
			"title": "爱小圈",
			"href":  "http://www.ixiaoquan.com",
		},
		{
			"title": "Github",
			"href":  "https://github.com/quarkcms",
		},
	},
}
