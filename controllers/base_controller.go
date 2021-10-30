package controllers

import (
	"AssetInventory/models"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	User    models.User
}

/**
判断是否登录
*/
func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginUser")
	if loginuser != nil {
		this.IsLogin = true
		this.User = loginuser.(models.User)

	} else {
		this.IsLogin = false
	}
}
