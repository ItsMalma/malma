package controller

import (
	"malma/pkg/response"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type AssetsController struct {
}

func NewAssetsController() *AssetsController {
	assetsController := new(AssetsController)

	return assetsController
}

func (c *AssetsController) Assets(ctx *gin.Context) {
	res, contentType, err := response.Assets(strings.Split(ctx.Param("filePath"), "/")...)
	if err != nil {
		if os.IsNotExist(err) {
			ctx.String(404, "Not Found")
		} else {
			ctx.String(500, "Internal Server Error")
		}
		return
	}
	ctx.Data(200, contentType, res)
}

func (c *AssetsController) Routing(router gin.IRouter) {
	router.GET("/assets/*filePath", c.Assets)
}
