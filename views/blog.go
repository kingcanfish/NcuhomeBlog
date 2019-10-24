package views

import (
	"NcuhomeBlog/lib"
	"NcuhomeBlog/model"
	"NcuhomeBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	g.GET("/:id", func(context *gin.Context) {
		result, err := GetBlogByID(context)
		if err != nil {
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
	blog:= new(model.BlogModel)
	err := lib.GetDB().Table(new(model.BlogModel).TableName()).Where("id = ?", id).Find(&blog)
	if err!=nil {
		return utils.FmtErrorReturn(err)
	} else {
		return utils.FmtNormalReturn(blog)
	}
}