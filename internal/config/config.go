package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error reading .env file. Try to get from environment instead.")
		for _, e := range os.Environ() {
			split := strings.Split(e, "=")
			k := split[0]
			_ = viper.BindEnv(k)
		}
		viper.AutomaticEnv()
	}

	initLogConfig()
	initAppConfig()
	initServerConfig()
	initDatabaseConfig()
	//initDatabaseAppierConfig()
	//initCacheConfig()
}
