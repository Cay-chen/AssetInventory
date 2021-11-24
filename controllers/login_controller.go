package controllers

import (
	"AssetInventory/models"
	"AssetInventory/utils"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type pd_name struct {
	Id               int
	PersonNo         int
	PersonName       string
	PersonCreateTime string
	PersonPsd        string
	PersonType       int
	PersonNum        string
}

type LoginController struct {
	BaseController
}

const (
	Date     = "2006-01-02"
	DateTime = "2006-01-02 15:04:05"
)

func (c *LoginController) Get() {

	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["error"]; ok {
		c.Data["Err"] = "inline"
		c.Data["Err_msg"] = n
	} else {
		c.Data["Err"] = "none"
	}
	c.TplName = "login.html"
	/*user :=models.User{UserNo: 1, UserName: "测试", UserCreateTime: time.Now().Format(DateTime), UserPsd: utils.Md5String32("123"), UserType: 1, UserNum: "JY555"}
	_, err := models.InsertUser(user)
	if err != nil {
		logs.Debug("注册失败！")
		return
	}else {
		logs.Debug("注册成功！")
	}*/
}
func (c *LoginController) Post() {
	userName := c.GetString("UserName")
	password := c.GetString("password")
	logs.Debug("username:" + userName + "    password" + password)
	row := models.QueryUserWithUser(userName, utils.Md5String32(password))
	user := new(models.User)
	err := row.Scan(&user.UserNo, &user.UserName, &user.UserType, &user.UserNum, &user.UserState)
	if err != nil {
		logs.Error(err)
		logs.Debug("用户名或者密码错误！")
		flash := beego.NewFlash()
		flash.Error("用户名或者密码错误！")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
		return
	}
	if user.UserNo > 0 {
		loginUser, _ := json.Marshal(user)
		err := c.SetSession("loginUser", string(loginUser))
		if err != nil {
			logs.Debug("Session存储出错！")
			return
		}
		logs.Debug("登录成功！")

		c.Redirect("/item", 302)
	} else {
		logs.Debug("用户名或者密码错误！")
		flash := beego.NewFlash()
		flash.Error("用户名或者密码错误！")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
	}

}
