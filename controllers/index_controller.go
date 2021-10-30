package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	if c.IsLogin {
		c.Data["UserName"] = c.User.UserName
		c.TplName = "index.html"

	} else {
		c.Redirect("/", 302)
	}

}
