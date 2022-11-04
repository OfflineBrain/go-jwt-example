package main

import (
	"github.com/gin-gonic/gin"
	"github.com/offlinebrain/go-jwt-example/auth"
	"github.com/offlinebrain/go-jwt-example/controller"
	"github.com/offlinebrain/go-jwt-example/repository/inmem"
)

func main() {
	inmem.NewUserRepository()

	router := initRouter()
	_ = router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user", controller.RegisterUser)
		secured := api.Group("/secured").Use(auth.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}
