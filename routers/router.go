package routers

import (
	"beegodemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/add", &controllers.ArticleController{}, "Get:AddArticle")
	beego.Router("/article/edit", &controllers.ArticleController{}, "Get:EditArticle")
	beego.Router("/user", &controllers.UserController{})
	//beego.Router("/user/add", &controllers.UserController{}, "Get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{}, "Post:DoAddUser")
	beego.Router("/user/edit", &controllers.UserController{}, "Get:Edit")
	beego.Router("/user/doEdit", &controllers.UserController{}, "Post:EditUser")
	beego.Router("/user/getJson", &controllers.UserController{}, "Get:GetUser")

	beego.Router("/good", &controllers.GoodController{})
	beego.Router("/good/add", &controllers.GoodController{}, "POST:DoAdd")
	beego.Router("/good/edit", &controllers.GoodController{}, "PUT:DoEdit")
	beego.Router("/good/delete", &controllers.GoodController{}, "DELETE:DoDelete")

}
