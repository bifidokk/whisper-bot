package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("app")
}

func GetEnv(key string) string {
	value, ok := viper.Get(key).(string)

	if !ok {
		fmt.Printf("Invalid type assertion")
	}

	return value
}
