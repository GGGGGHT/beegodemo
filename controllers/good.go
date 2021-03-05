package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/beego/beego/core/config/xml"
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

	//fmt.Println(c.GetString("title"))
	c.Ctx.WriteString("do add: " + c.GetString("name"))
}

func (c *GoodController) DoEdit() {
	p := Product{}

	if err := c.ParseForm(&p); err != nil {
		c.Ctx.WriteString("get data error")
		return
	}

	fmt.Printf("%#v", &p)
	c.Ctx.WriteString("do edit" + c.GetString("name"))
}

func (c *GoodController) DoDelete() {
	//id, err := c.GetInt("id")
	id := c.GetString("id")
	//if err != nil {
	//	panic(err)
	//fmt.Println(err)
	//c.Ctx.WriteString("do delete err: id must be int" )
	//return
	//}
	fmt.Printf("%T -> %v\n", id, id)
	c.Ctx.WriteString("do delete where id = " + id)
}

type Product struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

func (c *GoodController) XML() {
	article := Article{}
	//xml.Unmarshal(c.Ctx.Input.RequestBody)
	//str := string(c.Ctx.Input.RequestBody)
	//beego.Info(str)
	//c.Ctx.WriteString("result: " + str)
	var err error
	if e := xml.Unmarshal(c.Ctx.Input.RequestBody, &article); e != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	fmt.Printf("%#v\n", article)
	c.Data["json"] = article
	c.ServeJSON()
}
