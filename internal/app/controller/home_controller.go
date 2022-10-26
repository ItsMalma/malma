package controller

import (
	"malma/pkg/response"

	"github.com/gin-gonic/gin"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	homeController := new(HomeController)

	return homeController
}

func (c *HomeController) index(ctx *gin.Context) {
	res, err := response.View("index.html")
	if err != nil {
		ctx.String(500, "Internal Server Error")
		return
	}
	ctx.Data(200, "text/html", res)
}

func (c *HomeController) Routing(router gin.IRouter) {
	router.GET("", c.index)
	router.GET("/home", c.index)
	router.GET("/index", c.index)
}
