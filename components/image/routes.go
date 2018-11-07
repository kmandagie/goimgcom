package image

import (
	"github.com/labstack/echo"
)

// Routes Collection All Routes in Image Package
func Routes(g *echo.Group) {
	g.Add("GET", "/image/optimize/*", Optimize)
}
