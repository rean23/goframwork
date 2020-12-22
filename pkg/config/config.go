package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Viper *viper.Viper

type StrMap map[string]interface{}

func init() {
	//初始化viper对象
	Viper = viper.New()

	//设置配置文件名
	Viper.SetConfigFile(".env")

	//配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	Viper.SetConfigType("env")

	//设置配置文件路径
	Viper.AddConfigPath(".")

	//读取配置文件，读不到报错
	if err := Viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	Viper.AutomaticEnv()
	
	
	Viper.Set("app", map[string]interface{}{
		"a":1,
		"b":2,
	});
}

func Get(Key string, defaultValue ...interface{}) interface{} {

	if !Viper.IsSet(Key) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(Key)
}
