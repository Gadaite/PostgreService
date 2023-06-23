package main

import (
	conf "PostgreService/config"
	_ "PostgreService/docs"
	"PostgreService/routes"
	"fmt"
)

// @title POSTGRESQL-SERVER-GOLANG
// @version 1.0
// @description 使用golang语言,整合Swagger实现简单Postgres整合Gis服务
// @host localhost:2345
// @BasePath
func main() {
	conf.InitPostGresDB()
	conf.GetRedisConnect("tcp", "192.168.1.10:6379", 1)
	//connect := conf.GetRedisConnect("tcp", "1.15.94.16:6380", 0)
	//keys := conf.GetRedisAllKeys(connect)
	//fmt.Println(keys)
	initRoutes := routes.InitRoutes()
	err := initRoutes.Run("localhost:2345")
	if err == nil {
		panic(err.Error())
	} else {
		fmt.Sprint("Run : localhost:2345")
	}
}
