package bootstrap

import (
	"goframwork/pkg/config"
	"goframwork/pkg/model"
	"time"
)

func ConnectDB() {
	db := model.Connect()

	sqlDB, _ := db.DB()

	// SetMaxIdleConns设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))

	// SetMaxOpenConns设置到数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))

	// SetConnMaxLifetime设置连接可以重用的最长时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds"))*time.Second)
}
