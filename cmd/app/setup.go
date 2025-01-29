package app

import (
	"awesomeProjectSamb/internal/config"
	"awesomeProjectSamb/internal/database"
)

func Init() {
	config.Init()

	database.InitDB()
}

func Shutdown() {
	for _, conn := range database.DBConnection {
		_ = conn.Close()
	}
}
