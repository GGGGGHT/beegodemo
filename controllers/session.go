package controllers

import (
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

func (c *SessionController) Get() {
	c.TplName = "session.html"
}
