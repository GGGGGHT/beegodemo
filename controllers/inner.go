package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type InnerController struct {
	beego.Controller
}

func (c *InnerController) Get() {
	now := time.Now()
	c.Data["now"] = now
	c.Data["title"] = "这是文章标题"
	userInfo := make(map[string]interface{})
	userInfo["username"] = "zs"
	userInfo["age"] = 20
	c.Data["m"] = userInfo
	c.Data["unix"] = 1587880013

	c.TplName = "innerAPI.html"
}
