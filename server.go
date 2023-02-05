package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/tumininucodes/gin-crash-course/controller"
	"github.com/tumininucodes/gin-crash-course/entity/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)


func main() {

	server := gin.Default()

	server.Static("/css", ".templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid"})
		}
		ctx.JSON(200, videoController.Save(ctx))
	})

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid"})
			}
			ctx.JSON(200, videoController.Save(ctx))
		})
	}

	// viewRoutes := server.Group("/view")
	// {
	// 	viewRoutes.GET("/videos", videoController)
	// }

	server.Run(":8080")
}
