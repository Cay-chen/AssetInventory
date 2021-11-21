package controllers

import "AssetInventory/models"

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	if c.IsLogin {
		depList, _ := models.GetDepList()
		c.Data["DepList"] = depList
		c.TplName = "index.html"
	} else {
		c.Redirect("/", 302)
	}

}
