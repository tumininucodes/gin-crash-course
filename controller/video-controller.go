package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tumininucodes/gin-crash-course/entity"
	"github.com/tumininucodes/gin-crash-course/entity/service"
	"github.com/tumininucodes/gin-crash-course/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindALl()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBind(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindALl()
	data := gin.H {
		"title": "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}