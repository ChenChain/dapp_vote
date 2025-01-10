package util

import "github.com/gin-gonic/gin"

func ReturnErrResp(c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(500, err.Error())
		return true
	}
	return false
}

func ReturnResp(c *gin.Context, data interface{}) {
	c.JSON(200, data)
}
