package model

import (
	"fmt"
	"goframwork/pkg/config"
)
import "github.com/go-sql-driver/mysql"
import gormysql "gorm.io/driver/mysql"
import "gorm.io/gorm"

func Connect() *gorm.DB {
	mysqlConfig := mysql.Config{
		User:                 config.GetString("database.mysql.username"),
		Passwd:               config.GetString("database.mysql.password"),
		Addr:                 fmt.Sprintf("%s:%s", config.GetString("database.mysql.host"), config.GetString("database.mysql.port")),
		Net:                  "tcp",
		DBName:               config.GetString("database.mysql.database"),
		Collation:            "utf8mb4_unicode_ci",
		AllowNativePasswords: true,
	}

	dsn := mysqlConfig.FormatDSN()

	db, _ := gorm.Open(gormysql.Open(dsn), &gorm.Config{})

	return db
}
