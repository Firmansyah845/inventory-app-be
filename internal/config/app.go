package config

import "github.com/spf13/viper"

type AppConfig struct {
	DocsPath string
	AppEnv   string
	Name     string
	Version  string
}

var App AppConfig

func initAppConfig() {
	viper.SetDefault("APP_ENV", "production")
	viper.SetDefault("APP_NAME", "ads-api-service")
	//App.DocsPath = mustGetString("DOCS_PATH")
	App.AppEnv = mustGetString("APP_ENV")
	App.Name = mustGetString("APP_NAME")
	App.Version = mustGetString("APP_VERSION")
}
