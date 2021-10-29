package utils

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

var db *sql.DB

func InitMysql()  {
	logs.Debug("数据库初始化……………")
	driveName,_:= beego.AppConfig.String("driveName")
	mysqlConn, _ := beego.AppConfig.String("mysqlConn")

	db1, err := sql.Open(driveName, mysqlConn)
	if err != nil {
		return
	}else {
		db =db1
	}


}
/**
操作数据库
 */
func ModifyDB(sql string,args ...interface{}) (int64,error) {
	result,err:=db.Exec(sql,args)
	if err!=nil{
		logs.Debug(err)
		return 0,err
	}
	count ,err :=result.RowsAffected()
	if err!=nil {
		logs.Debug(err)
		return 0,err

	}
	return count,nil

	
}
/*
查询
 */
func QueryRowDB(sql string,args ...interface{})  *sql.Row{
	return db.QueryRow(sql,args)
}