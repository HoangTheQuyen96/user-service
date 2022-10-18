package config

import (
	"github.com/spf13/viper"
)

func SetupViper() *viper.Viper {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("USER_SVC")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	v := viper.GetViper()

	return v
}

var Cfg = SetupViper()
