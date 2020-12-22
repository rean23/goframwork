package model

import "goframwork/pkg/config"

func connect() *gorm.DB {
	mysql.Config{
		User:   config.Get("username"),
		Passwd: config.Get("password"),
		Addr:   config.Get("host"),
		Net:    "tcp",
		DBName: config.Get("database"),
	}.FormatDSN()
}
