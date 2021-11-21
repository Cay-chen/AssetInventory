package controllers

import (
	"AssetInventory/models"
	"AssetInventory/utils"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"strconv"
)

type DataServletController struct {
	BaseController
}
type TableData struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

func (c *DataServletController) Get() {
	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		limit, _ := c.GetInt("limit")
		page, _ := c.GetInt("page")
		beginNo := (page - 1) * limit
		logs.Debug(page)
		logs.Debug(limit)
		switch id {
		case "users":
			data, _ := models.GetListUserInfo(strconv.Itoa(beginNo), strconv.Itoa(limit))
			result, _ := json.Marshal(GetTableData("成功！", data, 0, utils.QueryTableTotalCount("pd_user_info")))
			c.Ctx.WriteString(string(result))
			break
		case "cards":
			where := ""
			depName := c.GetString("depName")
			assetsNo := c.GetString("assetsNo")
			if depName != "" {
				where = where + "WHERE user_dep='" + depName + "'"
				if assetsNo != "" {
					where = where + " AND card_no='" + assetsNo + "'"
				}
			} else {
				if assetsNo != "" {
					where = where + "WHERE card_no='" + assetsNo + "'"
				}
			}
			data, _ := models.GetListCardInfo(where, strconv.Itoa(beginNo), strconv.Itoa(limit))
			result, _ := json.Marshal(GetTableData("成功！", data, 0, utils.QueryTableTotalCount("pd_card_info "+where)))
			c.Ctx.WriteString(string(result))
			break
		default:
			c.TplName = "403.html"
		}

	} else {
		c.TplName = "403.html"
	}

}

func GetTableData(msg string, data []interface{}, code, count int) TableData {
	tableDate := TableData{}
	tableDate.Data = data
	tableDate.Code = code
	tableDate.Msg = msg
	tableDate.Count = count
	return tableDate
}
