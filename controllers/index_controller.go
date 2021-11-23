package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	if c.IsLogin {
		c.TplName = "pd_item.html"
	} else {
		c.Redirect("/", 302)
	}

}
