package env

import (
	"github.com/spf13/viper"
)

// 设置值
func Set(key string, value interface{}) {
	viper.SetConfigFile(".env")
	viper.Set(key, value)
}

// 获取值
func Get(key ...string) interface{} {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		if len(key) == 2 {
			if viper.Get(key[0]) == nil {
				return key[1]
			}
		} else {
			return nil
		}
	}

	if len(key) == 2 {
		if viper.Get(key[0]) == nil {
			return key[1]
		}
	}

	return viper.Get(key[0])
}
