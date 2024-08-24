package config

import (
	"os"

	"github.com/spf13/viper"
)

func SetupConfig() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(wd)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func GetTitle() (string, error) {
	title := viper.GetString("title")
	return title, nil
}

func GetPort() (string, error) {
	port := viper.GetString("port")
	return ":" + port, nil
}
