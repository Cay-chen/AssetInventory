package controllers

import (
	"AssetInventory/models"
	"github.com/beego/beego/v2/core/logs"
)

type DateInsertController struct {
	BaseController
}

func (c *DateInsertController) Get() {

	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "item":
			PdMsg := c.GetString("PdMsg")
			PdStartTime := c.GetString("PdStartTime")
			PdEndTime := c.GetString("PdEndTime")
			PdCreatName := c.User.UserNum
			ok, err := models.InsertItem(PdMsg, PdStartTime, PdCreatName, PdEndTime)
			if err != nil {
				logs.Debug(err)
				Res := models.BackRes("false", "创建失败！")
				c.Ctx.WriteString(Res)
				return
			}
			if ok > 0 {
				Res := models.BackRes("true", "创建成功！")
				c.Ctx.WriteString(Res)
			} else {
				Res := models.BackRes("false", "创建失败！")
				c.Ctx.WriteString(Res)
			}
			break
		}
	} else {
		Res := models.BackRes("false", "你没有权限！")
		c.Ctx.WriteString(Res)
	}
}
