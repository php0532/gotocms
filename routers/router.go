package routers

import (
	"github.com/astaxie/beego"
	"github.com/php0532/gotocms/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
