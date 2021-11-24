package models

import (
	_ "github.com/go-sql-driver/mysql"
)

func BackRes(result, msg string) string {
	return "{\"result\":" + result + ",\"msg\":\"" + msg + "\"}"
}
