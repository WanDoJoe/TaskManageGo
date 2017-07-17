package routers

import (
	"taskmanage/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.UserController{},"*:Login")
	beego.Router("/adduser", &controllers.UserController{},"*:AddUser")
}
