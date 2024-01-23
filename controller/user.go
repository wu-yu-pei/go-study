package controller

import (
	"github.com/gin-gonic/gin"
	"go-server-demo/service"
	"go-server-demo/tools"
)

type UserController struct{}

func (uc *UserController) Router(engine *gin.Engine) {
	engine.GET("/user", findUser)
	engine.GET("/user/err", errUser)
}

func findUser(ctx *gin.Context) {
	result := service.FindUser("wuyupei")
	tools.Success(ctx, result)
}

func errUser(ctx *gin.Context) {
	result := service.ErrUser("wuyupei")
	tools.Fail(ctx, result)
}
