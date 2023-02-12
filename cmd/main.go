package main

import (
	_ "github.com/PlanVX/aweme/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title aweme
// @version 1.0
// @description aweme api
// @contact.name PlanVX
// @contact.url https://github.com/PlanVX
// @license.name Apache 2.0
// @license.url https://github.com/PlanVX/aweme/blob/main/LICENSE
// @host localhost:8080
// @BasePath /douyin
func main() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
