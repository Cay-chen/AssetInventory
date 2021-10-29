package main

import (
	_ "AssetInventory/models"
	_ "AssetInventory/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

