package database

import (
	"database/sql"
	"log"

	"awesomeProjectSamb/internal/config"
	"awesomeProjectSamb/pkg/db"
)

const (
	MysqlDB = "mysqlDB"
)

var DBConnection = map[string]*sql.DB{}

func InitDB() {
	dbConfig := db.Config{
		Driver:          "mysql",
		URL:             config.Database.ConnectionURL(),
		MaxIdleConns:    config.Database.MaxPoolSize,
		MaxOpenConns:    config.Database.MaxPoolSize,
		ConnMaxLifeTime: config.Database.ConnectionMaxLifeTime,
	}

	mysqlDB, err := db.NewDB(&dbConfig)
	if err != nil {
		log.Fatalln("failed to initialize MySQL DB: " + config.Database.ConnectionURL() + " : " + err.Error())
	}
	DBConnection[MysqlDB] = mysqlDB
}
