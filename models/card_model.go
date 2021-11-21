package models

import (
	"AssetInventory/utils"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

type CardInfo struct {
	CardNo  string `json:"card_no"`
	Name    string `json:"name"`
	Spec    string `json:"spec"`
	Place   string `json:"place"`
	UserDep string `json:"user_dep"`
	User    string `json:"user"`
}

func GetListCardInfo(where, beginNo, pageSize string) ([]interface{}, error) {
	sqlS := fmt.Sprintf("SELECT ifnull(card_no,' ') as card_no,ifnull(name,' ') as name,ifnull(spec,' ') as spec,ifnull(place,' ') as place ,ifnull(user_dep,' ') as user_dep,ifnull(user,' ') as user FROM pd_card_info %s LIMIT %s,%s", where, beginNo, pageSize)
	rows, err := utils.QueryDB(sqlS)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var artList []interface{}
	for rows.Next() {
		cardNo := ""
		name := ""
		spec := ""
		place := ""
		userDep := ""
		user := ""
		rows.Scan(&cardNo, &name, &spec, &place, &userDep, &user)
		art := CardInfo{cardNo, name, spec, place, userDep, user}
		artList = append(artList, art)
	}
	return artList, nil

}
