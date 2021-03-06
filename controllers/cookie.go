package controllers

import (
	"github.com/astaxie/beego"
)

type CookieController struct {
	beego.Controller
}

func (c *CookieController) Get() {
	c.Ctx.SetCookie("username", "zhangsan")
	c.Ctx.SetCookie("password", "lisi")
	c.TplName = "index.tpl"
}
