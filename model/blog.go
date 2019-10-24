package model

import (
	"NcuhomeBlog/lib"
	"time"
)

type BlogModel struct {
	ID int
	Title string
	Author string
	Type_ string
	CreateTime time.Time
	ContentMD string
	ContentHTML string
}

func (model *BlogModel)TableName() string  {
	return "ncuhome_blog"

}

func (model *BlogModel)CreateTable()  {
	createTable(model.TableName(), model)
}

func (model *BlogModel) CheckExists() (bool, error)  {
	if model.Title!="" {
		return model.CheckExistsByTitle()
	}
}


func (model *BlogModel) CheckExistsByTitle() (bool, error)  {
	blog:= &BlogModel{Title:model.Title}
	has,err := lib.GetDB().Get(blog)
	if err!=nil {
		return has,err

	}else {
		return has ,nil
	}
}

func (model *BlogModel) CheckExistsByID() (bool,error) {

}
