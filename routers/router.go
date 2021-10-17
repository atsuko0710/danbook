package routers

import (
	"danbook/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // beego.Router("/", &controllers.MainController{})

	// 首页&分类&详情
	beego.Router("/", &controllers.HomeController{}, "get:Index")
	beego.Router("/category", &controllers.HomeController{}, "get:Category")

	// 图书
	beego.Router("/book/:key", &controllers.BookController{}, "get:Detail")
}
