package model

import (
	"NcuhomeBlog/lib"
	"errors"
	"time"
)

type BlogModel struct {
	ID int  `xorm:"'id' int pk autoincr" json:"id"`
	Title string `xorm:"'title' varchar(256) notnull" json:"title" binding:"required"`
	Author string `xorm:"'author' varchar(64) notnull" json:"author" bingding:"required"`
	Type_ string `xorm:"'type' varchar(64) notnull" json:"type"`
	CreateTime time.Time`xorm:"'create_time' datetime notnull" json:"-"`
	ContentMD string `xorm:"'content_md' text" json:"content_md"`
	ContentHTML string `xorm:"'content_html' text not null" json:"content_html" binding:"required"`
}

func (model *BlogModel)TableName() string  {
	return "ncuhome_blog"

}

func (model *BlogModel)CreateTable()  {
	createTable(model.TableName(), model)
}

func (model *BlogModel) CheckExists() (bool, error)  {
	if model.Title!= "" {
		return model.CheckExistsByTitle()
	}
	return true, errors.New("cols error")
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

