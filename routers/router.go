package routers

import (
	"AssetInventory/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/index", &controllers.IndexController{})
	beego.Router("/exit", &controllers.ExitController{})
}
