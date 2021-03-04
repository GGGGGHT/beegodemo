package controllers

import (
	"github.com/astaxie/beego"
)

type GoodController struct {
	beego.Controller
}

func (c *GoodController) Get() {
	// c.Data["Website"] = "beego.me"

	c.Data["title"] = "Beego Test"
	c.Data["num"] = "1"
	c.TplName = "good.html"
}

func (c *GoodController) DoAdd() {
	c.Ctx.WriteString("do add")
}

func (c *GoodController) DoEdit() {
	c.Ctx.WriteString("do edit")
}

func (c *GoodController) DoDelete() {
	c.Ctx.WriteString("do delete")
}
