package models

import (
	"AssetInventory/utils"
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

type User struct {
	UserNo         int    `json:"UserNo"`
	UserName       string `json:"UserName"`
	UserCreateTime string `json:"UserCreateTime"`
	UserPsd        string `json:"UserPsd"`
	UserType       int    `json:"UserType"`
	UserNum        string `json:"UserNum"`
	UserState      int    `json:"UserState"`
}

// QueryUserWightCon 按条件查询/*
func QueryUserWightCon(con string) int {
	sqlS := fmt.Sprintf("SELECT UserNo AS id FROM pd_user_info %s", con)
	logs.Debug("USER Mysql语句:" + sqlS)
	row := utils.QueryRowDB(sqlS)
	id := 0
	err := row.Scan(&id)
	if err != nil {
		logs.Error(err)
		return 0
	}
	return id
}

// InsertUser 插入用户数据/*
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("INSERT INTO pd_user_info (UserName,UserNum,UserPsd,UserType,UserCreateTime,UserState) VALUES (?,?,?,?,?,?)", user.UserName, user.UserNum, user.UserPsd, user.UserType, user.UserCreateTime, "0")

}

// QueryUserWithParam 跟据用户名和密码查询ID/*
func QueryUserWithParam(username, password string) int {
	sqlS := fmt.Sprintf("WHERE UserNum = '%s' and UserPsd = '%s'", username, password)
	return QueryUserWightCon(sqlS)

}

// QueryUserWithUserName 根据用户名查询ID/*
func QueryUserWithUserName(username string) int {
	sqlS := fmt.Sprintf("WHERE UserNum ='%s'", username)
	return QueryUserWightCon(sqlS)

}

// QueryUserWithUser 根据用户和密码查询用户信息/*
func QueryUserWithUser(username, password string) *sql.Row {
	sqlS := fmt.Sprintf("SELECT UserNo,UserName,UserType,UserNum,UserState FROM pd_user_info WHERE UserNum = '%s' and UserPsd = '%s'", username, password)
	logs.Debug("Mysql语句:" + sqlS)
	return utils.QueryRowDB(sqlS)

}

func GetListUserInfo(beginNo, pageSize string) ([]interface{}, error) {
	rows, err := utils.QueryTableLimitData("pd_user_info", beginNo, pageSize)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var artList []interface{}
	for rows.Next() {
		UserNo := 0
		UserName := ""
		UserCreateTime := ""
		UserPsd := ""
		UserType := 0
		UserNum := ""
		UserState := 0
		rows.Scan(&UserNo, &UserName, &UserCreateTime, &UserPsd, &UserType, &UserNum, &UserState)
		art := User{UserNo, UserName, UserCreateTime, "", UserType, UserNum, UserState}
		artList = append(artList, art)
	}
	return artList, nil

}
