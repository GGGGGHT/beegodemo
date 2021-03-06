package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// 模板中绑定基本数据类型
	c.Data["website"] = "beego 学习"
	c.Data["title"] = "你好 beego"
	c.Data["num"] = 12
	c.Data["flag"] = false
	a := &Article{
		"111 title",
		"222 content",
	}
	c.Data["struct"] = a
	c.Data["sliceList"] = []string{"php", "java", "golang"}

	c.Data["structSlice"] = []Article{
		{"news 1", "content 1"},
		{"news 2", "content 2"},
		{"news 3", "content 3"},
	}

	c.Data["isLogin"] = true
	c.Data["isHome"] = false
	c.Data["isAbout"] = true

	c.Data["n1"] = 22
	c.Data["n2"] = 6
	// 获取app.conf中的配置信息
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	beego.Info("user: %s", user)
	beego.Info("pwd: %s", pwd)
	users := beego.AppConfig.Strings("mysqluser")
	beego.Info("users: %s", users)
	// 向AppConfig中设置值
	if err := beego.AppConfig.Set("set", "value"); err != nil {
		beego.Info("set value error")
	}

	beego.BConfig
	beego.Info("ok...")

	c.TplName = "index.tpl"
}

type Article struct {
	Title   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}
