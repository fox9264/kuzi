package controller

import (
	"demo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetQqList(c *gin.Context) {
	qq := c.Query("qq")
	qqList, err := models.FindByQq(qq)
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, qqList)
	}
}

func GetWeiBoList(c *gin.Context) {
	uid := c.Query("uid")
	weiBoList, err := models.FindByUid(uid)
	if err!= nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else {
		c.JSON(http.StatusOK, weiBoList)
	}
}