package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"time"
)

type pd_name struct {
	Id int
	PersonNo   int
	PersonName string
	PersonCreateTime string
	PersonPsd string
	PersonType int
	PersonNum string

}

type LoginController struct {
	beego.Controller
}

const (
	Date      = "2006-01-02"
	DateTime  = "2006-01-02 15:04:05"
)


func (c *LoginController) Get() {
	/*now := time.Now()
	nowDate := now.Format(InitDate)
	log.Info("当前日期：", nowDate)
	nowTime := now.Format(DateTime)
	log.Info("当前时间：", nowTime)*/



	now := time.Now()                  //获取当前时间
	/*timestamp := now.Unix()            //时间戳
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()  //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒*/
	fmt.Printf("当前时间："+now.Format(DateTime))



	c.Data["Err"] = "none"
	orm.RegisterModel(new(pd_name))
	o :=orm.NewOrm()
	abc :=new(pd_name)
	abc.PersonName="管理员"
	abc.PersonNum="JY412"
	abc.PersonPsd="123456"
	abc.PersonCreateTime= now.Format(DateTime)
	abc.PersonType=1
	abc.PersonNo=003
	//per :=Person{PersonNum: "JY123"}
	//res :=o.Read(&per)
	//var maps []orm.Params
	//var person Person
	fmt.Println(o.Insert(abc))

	/*err := o.Raw("select a.* from pd_name a where a.PersonNum =?", "JY123").QueryRow(&maps)
	if err != nil {
		return 
	}
	fmt.Println(maps)*/
	c.TplName = "login.html"
}
func (c *LoginController) Post() {
}
