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
			if ok > 0 {

			} else {

			}
			if err != nil {
				logs.Debug(err)

			}

		}

	}

}
