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
	has, err :=db.IsTableExist(name)
	if err!=nil {
		log.Fatalln("checkout table error:", err.Error())
	}else if !has {
		err = db.Sync2(source)
		if err !=nil {
			log.Fatalln("create table error:", name)
		}
	}

}
