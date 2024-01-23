package router

import (
	"github.com/gin-gonic/gin"
	"go-server-demo/controller"
)

func SetupRouter() *gin.Engine {
	app := gin.Default()

	new(controller.UserController).Router(app)

	return app
}
