package image

import "github.com/labstack/echo"

// Optimize will optimize automatically will return image value
func Optimize(c echo.Context) error {
	return c.String(200, c.ParamValues()[0])
}
