package main

import (
	oimage "github.com/kmandagie/goimgcom/components/image"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	v1 := e.Group("/api/v1")
	oimage.Routes(v1)
	e.Logger.Fatal(e.Start(":6969"))
}
