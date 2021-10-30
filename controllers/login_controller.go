package controllers

import (
	"AssetInventory/models"
	"AssetInventory/utils"
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
	/*
		id:=models.QueryUserWithParam(userName,utils.Md5String32(password))
		logs.Debug("查询ID；"+strconv.Itoa(id))
		if id>0 {
			logs.Debug("登录成功！")
			c.Redirect("/index",302)
		}else {
			logs.Debug("用户名或者密码错误！")
			flash :=beego.NewFlash()
			flash.Error("用户名或者密码错误！")
			flash.Store(&c.Controller)
			c.Redirect("/",302)
		}*/

	row := utils.QueryRowDB("SELECT * FROM pd_user_info WHERE UserNum = '?' and UserPsd = '?'", userName, password)

	//row :=models.QueryUserWithUser(userName,utils.Md5String32(password))
	/*UserNo :=0
	UserName:=""
	UserCreateTime:=""
	UserPsd:=""
	UserType:=0
	UserNum:=""
	UserState:=0*/
	logs.Debug("AAAA")
	user := new(models.User)

	err := row.Scan(&user.UserNo, &user.UserName, &user.UserCreateTime, &user.UserPsd, &user.UserType, &user.UserNum, &user.UserState)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Debug("bbbb！")

	if user.UserNo > 0 {
		//user :=new (models.User)
		//user :=models.User{UserNo: UserNo, UserName: UserName, UserCreateTime: UserCreateTime, UserPsd: UserPsd, UserType: UserType, UserNum: UserNum, UserState: UserState}
		err := c.SetSession("loginUser", user)
		if err != nil {
			logs.Debug("Session存储出错！")
			return
		}
		logs.Debug("登录成功！")

		c.Redirect("/index", 302)
	} else {
		logs.Debug("用户名或者密码错误！")
		flash := beego.NewFlash()
		flash.Error("用户名或者密码错误！")
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
	}

}
