package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	// c.Data["Website"] = "beego.me"

	/*c.Data["title"] = "Beego Test"
	c.Data["num"] = "1"
	c.TplName = "good.html"*/
	c.Ctx.WriteString("用户中心")
}

func (c *UserController) User() {
	// c.Data["Website"] = "beego.me"

	/*c.Data["title"] = "Beego Test"
	c.Data["num"] = "1"
	c.TplName = "good.html"*/
	//c.Ctx.WriteString("用户中心")
	c.TplName = "user.html"
}

func (c *UserController) DoAddUser() {
	uname := c.GetString("username")
	pwd := c.GetString("password")
	if id, err := c.GetInt("id"); err != nil {
		c.Ctx.WriteString("id must be int type!")
	} else {
		fmt.Printf("%v --- %T\n", id, id)
	}
	hobby := c.GetStrings("hobby")
	/*for t := range hobby {
	}*/
	fmt.Printf("hobby: %v ->  %T\n", hobby, hobby)
	fmt.Printf("name: %s,pwd : %s\n", uname, pwd)
	c.Ctx.WriteString("name: " + uname + ",pwd: " + pwd)
	//c.ParseForm()
}

func (c *UserController) Edit() {
	// c.Data["Website"] = "beego.me"

	/*c.Data["title"] = "Beego Test"
	c.Data["num"] = "1"
	c.TplName = "good.html"*/
	//c.Ctx.WriteString("用户中心")
	c.TplName = "user.html"
}

func (c *UserController) EditUser() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		c.Ctx.WriteString("error")
		return
	}

	fmt.Printf("%#v\n", u)
	c.Ctx.WriteString("parse user success.")
}

type User struct {
	Username string   `form:"username" json:"uname"`
	Password string   `form:"password" json:"pwd"`
	Hobby    []string `form:"hobby" json:"h"`
}

func (c *UserController) GetUser() {
	u := User{
		Password: "123456",
		Username: "zhangsan",
		Hobby:    []string{"1", "2"},
	}

	c.Data["json"] = u
	c.ServeJSON()
}
