package main

import (
	"beegodemo/models"
	_ "beegodemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.TemplateLeft = "<<"
	//beego.BConfig.WebConfig.TemplateRight = ">>"
	beego.AddFuncMap("hi", hello)
	beego.AddFuncMap("unix2Date", models.Unix2Date)
	beego.Run()
}

func hello(in string) (out string) {
	out = in + " world"
	return
}
