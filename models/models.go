package models

import (
	_ "github.com/go-sql-driver/mysql"
)

/*func init() {

	orm.RegisterDataBase()

}*/
const (
	_OB_MYSQL_CONN = "root:root@tcp(222.209.81.188:3333)/pandian?charset=utf8"
	_MYSQL_DDRIVER = "mysql"
)

func init() {
	//mysqlConn, _ := beego.AppConfig.String("mysqlConn")
	/*orm.RegisterDriver(_MYSQL_DDRIVER, orm.DRMySQL)    //可以不加
	//	orm.RegisterDataBase("default", "mysql", "root:root@/databasename?charset=utf8")
	err := orm.RegisterDataBase("default", _MYSQL_DDRIVER, mysqlConn)
	if err != nil {
		logs.Debug("数据库连接出错！")
		logs.Debug(err.Error())

		return
	}else {
		logs.Debug("数据库连接成功！")

	}*/

	/*open, err := sql.Open(_MYSQL_DDRIVER, mysqlConn)
	if err != nil {
		return
	}else {
		db =open

	}*/

	//ORM 必须注册一个别名为default的数据库，作为默认使用。
}
