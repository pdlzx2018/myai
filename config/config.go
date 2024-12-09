package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}

	Redis struct {
		Host     string
		Port     int
		Password string
		DB       int
	}

	JWT struct {
		Secret string
		Expire int // token过期时间(小时)
	}
}

var GlobalConfig Config

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&GlobalConfig)
}
