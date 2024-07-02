package config

import "github.com/spf13/viper"

func GetConfig(key string) (bool, string) {
	var value string
	ok := true
	if viper.IsSet(key) {
		value = viper.GetString(key)
	} else {
		ok = false
	}
	return ok, value

}
