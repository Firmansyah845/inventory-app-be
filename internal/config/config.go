package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func Init() {

	if os.Getenv("ENVIRONMRNT") == "test" {
		viper.SetConfigName("test")
	} else {
		viper.SetConfigName("application")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("./../../configs")
	viper.AddConfigPath("./../../../configs")
	viper.AddConfigPath("./../../../../configs")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.AutomaticEnv()

	initLogConfig()
	initAppConfig()
	initServerConfig()
	initDatabaseConfig()
}
