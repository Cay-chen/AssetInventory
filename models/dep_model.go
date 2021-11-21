package models

import (
	"AssetInventory/utils"
	"github.com/beego/beego/v2/core/logs"
)

type DepInfo struct {
	UserDep string `json:"user_dep"`
	User    string `json:"user"`
	Total   int    `json:"total"`
}

func GetDepList() ([]DepInfo, error) {
	querySql := "select user_dep,user,count(*) from pd_card_info group by user_dep"
	rows, err := utils.QueryDB(querySql)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var artList []DepInfo
	for rows.Next() {
		userDep := ""
		user := ""
		total := 0
		rows.Scan(&userDep, &user, &total)
		art := DepInfo{userDep, user, total}
		artList = append(artList, art)
	}
	return artList, nil
}
