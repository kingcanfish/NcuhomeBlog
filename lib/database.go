package lib

import (
	"NcuhomeBlog/conf"
	"github.com/go-xorm/xorm"
	"log"
)

import _ "github.com/go-sql-driver/mysql"

var db  *xorm.Engine

func InitMysql() {
	var err error

	db, err = xorm.NewEngine("mysql", conf.GetConfig().DatabaseURI)
	if err !=nil {
		log.Fatalln("connect_mysql" , err.Error())
	}
}

func GetDB() *xorm.Engine {
	return db
}

func init()  {
	InitMysql()
}


