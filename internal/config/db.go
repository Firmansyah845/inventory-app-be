package config

import (
	"fmt"
	"time"
)

type DatabaseConfig struct {
	Name                  string
	Host                  string
	User                  string
	Password              string
	Port                  int
	MaxPoolSize           int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ConnectionMaxLifeTime time.Duration
	SSLMode               string
}

var Database DatabaseConfig

func initDatabaseConfig() {

	Database = DatabaseConfig{
		Name:                  mustGetString("DB_NAME"),
		Host:                  mustGetString("DB_HOST"),
		User:                  mustGetString("DB_USER"),
		Password:              mustGetString("DB_PASSWORD"),
		Port:                  mustGetInt("DB_PORT"),
		MaxPoolSize:           mustGetInt("DB_POOL_SIZE"),
		ReadTimeout:           mustGetDurationMs("DB_READ_TIMEOUT"),
		WriteTimeout:          mustGetDurationMs("DB_WRITE_TIMEOUT"),
		ConnectionMaxLifeTime: mustGetDurationMinute("DB_CONNECTION_MAX_LIFETIME_MINUTE"),
		SSLMode:               mustGetString("SSL_MODE"),
	}
}

func (dc DatabaseConfig) ConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dc.User, dc.Password, dc.Host, dc.Port, dc.Name)
}

type DatabaseAppierConfig struct {
	Name                  string
	Host                  string
	User                  string
	Password              string
	Port                  int
	MaxPoolSize           int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ConnectionMaxLifeTime time.Duration
	SSLMode               string
}

var DatabaseAppier DatabaseAppierConfig

func initDatabaseAppierConfig() {

	DatabaseAppier = DatabaseAppierConfig{
		Name:                  mustGetString("DB_NAME_APPIER"),
		Host:                  mustGetString("DB_HOST_APPIER"),
		User:                  mustGetString("DB_USER_APPIER"),
		Password:              mustGetString("DB_PASSWORD_APPIER"),
		Port:                  mustGetInt("DB_PORT_APPIER"),
		MaxPoolSize:           mustGetInt("DB_POOL_SIZE_APPIER"),
		ReadTimeout:           mustGetDurationMs("DB_READ_TIMEOUT_APPIER"),
		WriteTimeout:          mustGetDurationMs("DB_WRITE_TIMEOUT_APPIER"),
		ConnectionMaxLifeTime: mustGetDurationMinute("DB_CONNECTION_MAX_LIFETIME_MINUTE_APPIER"),
		SSLMode:               mustGetString("SSL_MODE_APPIER"),
	}
}

func (dac DatabaseAppierConfig) ConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dac.User, dac.Password, dac.Host, dac.Port, dac.Name)
}
