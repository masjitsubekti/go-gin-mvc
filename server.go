package main

import (
	"github.com/gin-gonic/gin"
	"github.com/masjitsubekti/go-gin-mvc/controller"
	"github.com/masjitsubekti/go-gin-mvc/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})

	server.Run(":8080")
}
