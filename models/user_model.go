package models

import (
	"AssetInventory/utils"
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

type User struct {
	UserNo         int
	UserName       string
	UserCreateTime string
	UserPsd        string
	UserType       int
	UserNum        string
	UserState      int
}

/*
按条件查询
*/
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("SELECT UserNo AS id FROM pd_user_info %s", con)
	logs.Debug("USER Mysql语句:" + sql)
	row := utils.QueryRowDB(sql)
	id := 0
	err := row.Scan(&id)
	if err != nil {
		logs.Error(err)
		return 0
	}
	return id
}

/*
插入用户数据
*/
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("INSERT INTO pd_user_info (UserName,UserNum,UserPsd,UserType,UserCreateTime,UserState) VALUES (?,?,?,?,?,?)", user.UserName, user.UserNum, user.UserPsd, user.UserType, user.UserCreateTime, "0")

}

/*
跟据用户名和密码查询ID
*/
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("WHERE UserNum = '%s' and UserPsd = '%s'", username, password)
	return QueryUserWightCon(sql)

}

/*
根据用户名查询ID
*/
func QueryUserWithUserName(username string) int {
	sql := fmt.Sprintf("WHERE UserNum ='%s'", username)
	return QueryUserWightCon(sql)

}

/*
根据用户和密码查询用户信息
*/
func QueryUserWithUser(username, password string) *sql.Row {
	sql := fmt.Sprintf("SELECT * FROM pd_user_info WHERE UserNum = '%s' and UserPsd = '%s'", username, password)
	logs.Debug("Mysql语句:" + sql)
	return utils.QueryRowDB(sql)

}
