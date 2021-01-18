package config

import "goframwork/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		"env":   config.Env("APP_ENV", "production"),
		"debug": config.Env("APP_DEBUG", false),
		"port":  config.Env("APP_PORT", "3000"),
	})
}
