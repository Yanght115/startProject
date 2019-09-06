package main

import (
	"plan/pkg/config"
	"plan/pkg/gmysql"
	"plan/pkg/logging"
)

func init() {
	config.LoadConfig("conf")
	logging.InitLogger("")
	logging.Logger.Infof("init config and logger successfully, Environment is : %s", config.GlobalConfig.Env)
	gmysql.InitMySQLClient()
}

func main() {
	print("test")
	r := router.GetRouter()
}
