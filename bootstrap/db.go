package bootstrap

import (
	"fmt"
	"goframwork/pkg/model"
	"time"
)

func ConnectDB() {
	db := model.Connect()

	sqlDB, _ := db.DB()

	// SetMaxIdleConns设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns设置到数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime设置连接可以重用的最长时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println(sqlDB)
}
