package views

import (
	"NcuhomeBlog/lib"
	"NcuhomeBlog/model"
	"NcuhomeBlog/utils"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const author string = "NCUHOME"

func BlogRouteGroup(g *gin.RouterGroup)  {
	g.GET("/all", func(context *gin.Context) {
		result, err := GetAll(context)
		if err!=nil {
			context.JSON(http.StatusInternalServerError, result)
		}else {
			context.JSON(http.StatusOK, result)
		}

	})

	g.GET("", func(context *gin.Context) {
		result, err:= GetBlogListByType(context)
		if err!=nil {
			context.JSON(http.StatusInternalServerError, result)
		} else {
			context.JSON(http.StatusOK, result)
		}

	})
	g.GET("/get/:id", func(context *gin.Context) {
		result, err := GetBlogByID(context)
		if err != nil {
			context.JSON(http.StatusInternalServerError, result)
		} else {
			context.JSON(http.StatusOK, result)
		}
	})

	g.POST("/publish", func(context *gin.Context) {
		 if result, err:= CreateBlog(context); err!=nil {
		 	context.JSON(http.StatusInternalServerError, result)
		 } else {
			 context.JSON(http.StatusOK, result)
		 }


	})
}


//获取所有的文章
func GetAll(c *gin.Context)(map[string] interface{}, error)  {
	blogs:= make([]*model.BlogModel, 0)
	err:=lib.GetDB().Table(new(model.BlogModel)).Find(&blogs)
	if err!=nil {
		return utils.FmtErrorReturn(err)
	}
	return utils.FmtNormalReturn(blogs, "ok")


}

func GetBlogListByType(c *gin.Context) (map[string]interface{}, error)  {
	type_ := c.Query("type")
	blogs := make([]*model.BlogModel, 0)


	err:= lib.GetDB().Table(new(model.BlogModel).TableName()).Where("type = ?", type_).Find(&blogs)
	if err!=nil {
		return utils.FmtErrorReturn(err)
	} else {
		return utils.FmtNormalReturn(blogs)
	}
}

func GetBlogByID(c *gin.Context) (map[string] interface{}, error) {
	id:=c.Param("id")
	if id == "0" {
		return utils.FmtErrorReturn(errors.New("id cant be 0"))
	}
	blogs:= make([] *model.BlogModel, 0)
	err := lib.GetDB().Table(new(model.BlogModel).TableName()).Where("id = ?", id).Find(&blogs)
	if err!=nil {
		return utils.FmtErrorReturn(err)
	} else {
		return utils.FmtNormalReturn(blogs[0])
	}
}

func CreateBlog(c *gin.Context) (map[string] interface{}, error) {
	modelBlog :=new(model.BlogModel)
	if err:= c.ShouldBindJSON(modelBlog); err!=nil {
		return utils.FmtErrorReturn(err)
	} else {
		 if has ,err := modelBlog.CheckExists(); err!=nil {
		 	return utils.FmtErrorReturn(err)
		 } else  if has{
			return utils.FmtNormalReturn("", "该记录已经存在！")
		 } else {
		 	session := lib.GetDB().NewSession()
		 	defer session.Close()
			 if err := session.Begin(); err!=nil {
				 return utils.FmtErrorReturn(err)
			 }
			 modelBlog.CreateTime = time.Now()
			 modelBlog.Author = author
			 if _ , err:= session.Insert(modelBlog); err!= nil {
			 	_ = session.Rollback()
			 	return utils.FmtErrorReturn(err)
			 }
			 _ = session.Commit()
			 return utils.FmtNormalReturn("", "提交成功～")
		 }

	}

}