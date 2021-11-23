package controllers

import "AssetInventory/models"

type AssetsController struct {
	BaseController
}

func (c *AssetsController) Get() {
	if c.IsLogin {
		depList, _ := models.GetDepList()
		c.Data["DepList"] = depList
		c.TplName = "index.html"
	} else {
		c.Redirect("/", 302)
	}

}
