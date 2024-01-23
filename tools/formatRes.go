package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = 200
	FAIL    int = 500
)

type ResponseContent = map[string]interface{}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseContent{
		"code": SUCCESS,
		"msg":  "ok",
		"data": data,
	})
}

func Fail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseContent{
		"code": FAIL,
		"msg":  "error",
		"data": data,
	})
}
