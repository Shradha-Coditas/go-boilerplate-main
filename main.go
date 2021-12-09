package main

import (
	"boilerplate/config"
	MySQL "boilerplate/library/mysql"
	redis "boilerplate/library/redis"
	"boilerplate/util/localization"
	"boilerplate/util/logwrapper"
	"boilerplate/util/server"
	"strings"
)

func main() {
	configData := config.LoadEnv()

	Logger := logwrapper.NewLogger(configData.Server)

	localization.LoadBundle(configData.Server)

	Logger.Info(strings.Repeat("~", 50))

	// connect to MysSql server

	sqlErr := MySQL.NewConnection(configData.MySQL)
	if sqlErr != nil {
		Logger.Fatal("Error connecting to SQL : ", sqlErr)
	}

	redisConnErr := redis.NewConnection(configData.Redis)
	if redisConnErr != nil {
		Logger.Fatal("Error connecting to redis : ", redisConnErr)
	}

	// start GIN server
	server.StartServer(configData.Server)
}
