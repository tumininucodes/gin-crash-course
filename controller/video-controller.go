package controller

import (

	"github.com/gin-gonic/gin"
	"github.com/tumininucodes/gin-crash-course/entity"
	"github.com/tumininucodes/gin-crash-course/entity/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindALl()
}

func (c* controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.ShouldBind(&video)
	c.service.Save(video)
	return video
}