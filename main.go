package main

import (
	_ "AssetInventory/models"
	_ "AssetInventory/routers"
	"AssetInventory/utils"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	utils.InitMysql()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 600
	beego.Run()
}
