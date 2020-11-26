package routers

import (
	"demo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.GET("/qq", controller.GetQqList)
		v1Group.GET("/weibo", controller.GetWeiBoList)
	}
	return r
}
