package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/masjitsubekti/go-gin-mvc/model"
	"github.com/masjitsubekti/go-gin-mvc/service"
	"github.com/masjitsubekti/go-gin-mvc/validators"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

// Constructor
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video model.Video
	err := ctx.BindJSON(&video)
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

// Load HTML
func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
