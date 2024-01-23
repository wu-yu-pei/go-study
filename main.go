package main

import (
	"fmt"
	"go-server-demo/middleware"
	"go-server-demo/router"
)

func main() {

	app := router.SetupRouter()

	middleware.Setup(app)

	err := app.Run("0.0.0.0:9090")

	if err != nil {
		fmt.Println("server is erring...", err)
	}

	fmt.Println("server is starting...")

}
