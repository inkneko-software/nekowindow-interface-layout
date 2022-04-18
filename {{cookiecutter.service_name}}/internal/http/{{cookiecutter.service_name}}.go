package http

import (
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/data"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

// VideoFileSystemController is a http controller.
type VideoHttpController struct {
	d   *data.Data
	log *log.Helper
}

func NewVideoHttpController(d *data.Data, logger log.Logger) *VideoHttpController {
	return &VideoHttpController{d: d, log: log.NewHelper(log.With(logger, "package", "http"))}
}

func (controller *VideoHttpController) ExampleHandler(ctx *gin.Context) {
	//Implement your handler here.
}
