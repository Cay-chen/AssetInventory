package controllers

type ItemController struct {
	BaseController
}

func (c *ItemController) Get() {
	if c.IsLogin {
		c.TplName = "pd_item.html"
	} else {
		c.Redirect("/", 302)
	}
}
