package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/main-shopping-ms/cfg"
	"github.com/prakash-p-3121/main-shopping-ms/controller/purchase_controller"
	"github.com/prakash-p-3121/main-shopping-ms/controller/signup_controller"
	"github.com/prakash-p-3121/restlib"
)

func main() {
	msConnectionsMap, err := restlib.CreateMsConnectionCfg("conf/microservice.toml")
	if err != nil {
		panic(err)
	}
	cfg.SetMsConnectionsMap(msConnectionsMap)

	router := gin.Default()
	publicGroup := router.Group("/shopping/public")
	publicGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	publicGroup.POST("/v1/signup", signup_controller.Signup)

	privateGroup := router.Group("/shopping/private")
	privateGroup.POST("/v1/purchase", purchase_controller.CreatePurchase)

	err = router.Run("127.0.0.1:3050")
	if err != nil {
		panic("Error Starting main-shopping-ms")
	}
}
