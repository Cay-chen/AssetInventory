package controllers

import "AssetInventory/models"

type CardController struct {
	BaseController
}

func (c *CardController) Get() {
	if c.IsLogin {
		depList, _ := models.GetDepList()
		c.Data["DepList"] = depList
		c.TplName = "pd_card_info.html"
	} else {
		c.Redirect("/", 302)
	}
}
