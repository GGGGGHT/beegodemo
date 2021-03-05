package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {
	name := c.GetString("username")
	password := c.GetString("password")
	fmt.Println("username: " + name + ", password: " + password + " login success!")

	c.Redirect("/", 302)
}
