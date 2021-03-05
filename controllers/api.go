package controllers

import (
	"github.com/astaxie/beego"
)

type APIController struct {
	beego.Controller
}

func (c *APIController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api interface: " + id)
}
