package models

import (
	"AssetInventory/utils"
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
	rows, err := utils.QueryTableLimitData("pd_card_info "+where, beginNo, pageSize)
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
		logs.Debug(art)
		artList = append(artList, art)
	}
	return artList, nil

}
