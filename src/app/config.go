package app

import (
	"os"

	"github.com/spf13/viper"
)

func InitConfig() error {
	var configFile = os.Getenv("APP_ENV_FILE")
	if os.Getenv("APP_ENV_FILE") == "" {
		configFile = ".env"
	}

	viper.SetConfigFile(configFile)
	viper.AutomaticEnv()

	var err = viper.ReadInConfig()

	return err
}
