package controllers

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	if c.IsLogin {
		c.TplName = "user_manage.html"
	} else {
		c.TplName = "403.html"
	}

}
