package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MysqlSetting ...
type MysqlSetting interface {
	GetUserName() string
	GetPassword() string
	GetHost() string
	GetDatabase() string
}

var (
	db *sql.DB
)

// NewDB 取得DB
func NewDB(setting MysqlSetting) *sql.DB {
	if db == nil {
		var err error
		datasource := fmt.Sprintf("%v:%v@tcp(%v)/%v", setting.GetUserName(), setting.GetPassword(), setting.GetHost(), setting.GetDatabase())
		db, err = sql.Open("mysql", datasource)
		if err != nil {
			log.Fatal(err)
		}
	}
	return db
}
