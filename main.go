package main

import (
	"github.com/kmandagie/goimgcom/components/css"
	oimage "github.com/kmandagie/goimgcom/components/image"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	v1 := e.Group("/api/v1")
	oimage.Routes(v1)
	css.Routes(v1)
	e.Logger.Fatal(e.Start(":6969"))
}
