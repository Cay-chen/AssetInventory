package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.TplName = "index.html"

	/*if c.IsLogin{
		c.TplName="index.html"

	}else {
		c.TplName="login.html"
	}*/

}
