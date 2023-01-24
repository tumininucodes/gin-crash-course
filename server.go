package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tumininucodes/gin-crash-course/controller"
	"github.com/tumininucodes/gin-crash-course/entity/service"
	"github.com/tumininucodes/gin-crash-course/middlewares"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOuput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOuput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run()
}