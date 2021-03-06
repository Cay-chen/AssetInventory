package routers

import (
	"AssetInventory/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	//beego.Router("/assets", &controllers.AssetsController{})
	beego.Router("/item", &controllers.ItemController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/card", &controllers.CardController{})
	beego.Router("/data/table/:id", &controllers.DataTableServletController{})
	beego.Router("/data/insert/:id", &controllers.DateInsertController{})

}
