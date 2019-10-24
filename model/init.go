package model

import (
	"NcuhomeBlog/lib"
	"log"
)

func init()  {
	blog := new(BlogModel)
	blog.CreateTable()
	
}

func createTable(name string, source interface{})  {
	db:=lib.GetDB()
	has, err :=db.Exist(name)
	if err!=nil {
		log.Fatalln("checkout table error:", name)
	}else if !has {
		err = db.Sync2(source)
		if err !=nil {
			log.Fatalln("create table error:", name)
		}
	}

}
