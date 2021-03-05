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

	c.TplName = "index.tpl"
}

type Article struct {
	Title   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}
