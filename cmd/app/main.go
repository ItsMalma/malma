package main

import (
	"fmt"
	"malma/internal/app/controller"

	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
)

type EnvConfig struct {
	Port int `env:"PORT" envDefault:"3000"`
}

func main() {
	envConfig := EnvConfig{}
	if err := env.Parse(&envConfig); err != nil {
		panic(err)
	}

	app := gin.Default()

	assesController := controller.NewAssetsController()
	homeController := controller.NewHomeController()

	assesController.Routing(app)
	homeController.Routing(app)

	if err := app.Run(fmt.Sprintf(":%v", envConfig.Port)); err != nil {
		panic(err)
	}
}
