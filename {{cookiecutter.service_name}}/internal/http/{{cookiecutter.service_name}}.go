package http

import (
	"nekowindow-backend/app/{{cookiecutter.kind}}/{{cookiecutter.department}}/{{cookiecutter.service_name}}/internal/data"
	"nekowindow-backend/pkg/net/http/middleware/response"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

// {{cookiecutter.serviceUpper}}FileSystemController is a http controller.
type {{cookiecutter.serviceUpper}}HttpController struct {
	d *data.Data
	log     *log.Helper
}

func New{{cookiecutter.serviceUpper}}HttpController(d *data.Data, logger log.Logger) *{{cookiecutter.serviceUpper}}HttpController {
	return &{{cookiecutter.serviceUpper}}HttpController{d:d, log: log.NewHelper(log.With(logger, "package", "http"))}
}

func (controller *{{cookiecutter.serviceUpper}}HttpController) ExampleHandler(ctx *gin.Context) {
	//Implement your handler here.
	req := v1.ExampleReq{}
	resp, err := service.ExampleService(ctx, req)
	if err != nil{
		response.AbortWithErrors(ctx, err)
	}
	response.Protobuf(ctx, resp)

}
