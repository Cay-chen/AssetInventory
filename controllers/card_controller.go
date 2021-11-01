package controllers

type CardController struct {
	BaseController
}

func (c *CardController) Get() {
	if c.IsLogin {
		c.TplName = "card_manage.html"
	} else {
		c.TplName = "403.html"
	}

}
