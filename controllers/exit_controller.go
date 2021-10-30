package controllers

type ExitController struct {
	BaseController
}

func (c *ExitController) Get() {
	if c.IsLogin {
		err := c.DelSession("loginUser")
		if err != nil {
			c.Redirect("/", 302)
			return
		} else {
			c.Redirect("/", 302)
		}
	} else {
		c.Redirect("/", 302)
	}

}
