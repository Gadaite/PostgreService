package config

import (
	"fmt"
	_ "github.com/go-spatial/geom"
	_ "github.com/go-spatial/tegola"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var PostDB *gorm.DB

var ShowConsoleResult bool = false

const (
	postgresUser     = "postgres"
	postgresPassword = "LYP809834049"
	postgresHost     = "192.168.1.10"
	postgresPort     = 5432
	postgresDbName   = "trajectory"
)

func InitPostGresDB() {
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDbName))
	db.LogMode(true)
	db.SingularTable(true)
	if err != nil {
		panic("connect to postgresql failed :" + err.Error())
	}
	PostDB = db
}
