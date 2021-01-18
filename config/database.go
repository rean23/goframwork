package config

import "goframwork/pkg/config"

func init() {
	config.Add("database", config.StrMap{
		"mysql": map[string]interface{}{
			"host":     config.Env("DB_HOST", "local"),
			"port":     config.Env("DB_PORT", "3306"),
			"username": config.Env("DB_USERNAME", "1"),
			"password": config.Env("DB_PASSWORD", "1"),
			"database": config.Env("DB_DATABASE", "1"),
			"charset":  config.Env("DB_CHARSET", "1"),

			"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 100),
			"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 50),
			"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
		},
	})
}
