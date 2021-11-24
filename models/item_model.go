package models

import (
	"AssetInventory/utils"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

type ItemBean struct {
	PdNo         string `json:"pd_no"`
	PdStartTime  string `json:"pd_start_time"`
	PdMsg        string `json:"pd_msg"`
	PdCreateName string `json:"pd_create_name"`
	PdStatus     string `json:"pd_status"`
	PdCreateTime string `json:"pd_create_time"`
	PdEndTime    string `json:"pd_end_time"`
}

func GetPdItemList(where, beginNo, pageSize string) ([]interface{}, error) {
	sqlS := fmt.Sprintf("select i.pd_no,i.pd_msg,u.UserName as pd_create_name,i.pd_start_time ,i.pd_end_time,i.pd_create_time,i.pd_status from pd_item i inner join pd_user_info u where i.pd_create_name =u.UserNum  %s order by i.pd_create_time desc limit %s,%s ", where, beginNo, pageSize)
	rows, err := utils.QueryDB(sqlS)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var artList []interface{}
	for rows.Next() {
		pdNo := ""
		pdStartTime := ""
		pdMsg := ""
		pdCreateName := ""
		pdStatus := ""
		pdCreateTime := ""
		pdEndTime := ""
		err := rows.Scan(&pdNo, &pdMsg, &pdCreateName, &pdStartTime, &pdEndTime, &pdCreateTime, &pdStatus)
		if err != nil {
			return nil, err
		}
		art := ItemBean{pdNo, pdStartTime, pdMsg, pdCreateName, pdStatus, pdCreateTime, pdEndTime}
		artList = append(artList, art)
	}
	return artList, nil
}

func InsertItem(PdMsg, PdStartTime, PdCreateName, IPdEndTime string) (int64, error) {
	return utils.ModifyDB("INSERT INTO pd_item (pd_msg,pd_start_time,pd_create_name,pd_create_time,pd_end_time) VALUES (?,?,?,NOW(),?)", PdMsg, PdStartTime, PdCreateName, IPdEndTime)
}
