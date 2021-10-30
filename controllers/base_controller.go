package controllers

import (
	"AssetInventory/models"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	User    *models.User
}

/**
判断是否登录
*/
func (c *BaseController) Prepare() {
	loginUser := c.GetSession("loginUser")
	if loginUser != nil {
		res := &models.User{}
		err := json.Unmarshal([]byte(loginUser.(string)), &res)
		if err != nil {
			c.IsLogin = false
			return
		} else {
			c.User = res
			c.IsLogin = true
		}

	} else {
		c.IsLogin = false
	}
}
