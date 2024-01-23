package middleware

import "github.com/gin-gonic/gin"

func Setup(engine *gin.Engine) {
	engine.Use(CORS())
}
