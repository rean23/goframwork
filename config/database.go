package config

import "goframwork/pkg/config"

func init() {
	config.Add("database", config.StrMap{
		"mysql": map[string]interface{}{
			"host":     config.Env("DB_HOST", "local"),
			"username": config.Env("DB_USERNAME", "1"),
			"password": config.Env("DB_PASSWORD", "1"),
			"database": config.Env("DB_DATABASE", "1"),
			"charset":  config.Env("DB_CHARSET", "1"),
		},
	})
}
