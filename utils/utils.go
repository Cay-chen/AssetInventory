package utils

import (
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {
	logs.Debug("数据库初始化……………")
	driveName, _ := beego.AppConfig.String("driveName")
	mysqlConn, _ := beego.AppConfig.String("mysqlConn")

	db1, err := sql.Open(driveName, mysqlConn)
	if err != nil {
		return
	} else {
		db = db1
	}

}

// ModifyDB /**
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	logs.Debug("操作数据库语句:" + sql)
	logs.Debug(args)
	result, err := db.Exec(sql, args...)
	if err != nil {
		logs.Debug("**********语句执行错误*******************")
		logs.Debug(err)
		logs.Debug("***************************************")

		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		logs.Debug(err)
		return 0, err

	}
	return count, nil

}

// QueryRowDB 查询单行数据/*
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// QueryDB 查询所有数据/*
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

// QueryTableLimitData 按照limit查询数据
func QueryTableLimitData(table, beginNo, pageSize string) (*sql.Rows, error) {
	sqlS := fmt.Sprintf("SELECT *  FROM %s LIMIT %s,%s", table, beginNo, pageSize)
	logs.Debug("MYSQL语句:" + sqlS)
	return QueryDB(sqlS)

}

// QueryTableTotalCount 查询表的总数量
func QueryTableTotalCount(table string) int {
	sqlS := fmt.Sprintf("SELECT COUNT(*) AS total FROM %s", table)
	logs.Debug("MYSQL语句:" + sqlS)
	row := QueryRowDB(sqlS)
	total := 0
	err := row.Scan(&total)
	if err != nil {
		return 0
	}
	return total
}
