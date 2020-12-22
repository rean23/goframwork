package config

import "goframwork/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		"env":   config.Env("APP_ENV", "local"),
		"debug": config.Env("APP_DEBUG", "1"),
	})
}
