package config

import (
	"fmt"
	"github.com/spf13/cast"
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

func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}

func Add(key string, configMap map[string]interface{}) {
	Viper.Set(key, configMap)
}

func Env(Key string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(Key, defaultValue[0])
	}
	return Get(Key)
}
