package controller

import (

	"github.com/gin-gonic/gin"
	"github.com/tumininucodes/gin-crash-course/entity"
	"github.com/tumininucodes/gin-crash-course/entity/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error 
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

func (c* controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBind(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}