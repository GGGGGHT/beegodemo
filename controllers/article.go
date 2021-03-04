package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	// c.Data["Website"] = "beego.me"

	//c.Data["title"] = "Beego Test"
	//c.Data["num"] = "1"
	//c.TplName = "good.html"
	c.Ctx.WriteString("你好 世界!")
}

func(c *ArticleController) AddArticle() {
	c.Ctx.WriteString("add article")
}

func(c *ArticleController) EditArticle() {
	id := c.GetString("id")
	//fmt.Printf("%v %T\n",id, id)
	if i, err := c.GetInt("id"); err != nil {
		beego.Info(err)
	} else {
		fmt.Printf("%v %T\n",i,i)
	}


	c.Ctx.WriteString("修改新闻:" + id)
}