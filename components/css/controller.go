package css

import (
	"strings"

	"github.com/kmandagie/goimgcom/components/utils"
	"github.com/labstack/echo"
)

// Optimize this function will return optimize css
func Optimize(c echo.Context) error {
	url := c.ParamValues()[0]
	css := utils.GetFile(url)
	buff := utils.ReaderToBuffer(css.Body)
	data := removeNewLine(string(buff))
	data = removeSpace(data)
	c.Response().Header().Set("Content-Type", "text/css")
	return c.String(200, data)
}

func removeNewLine(data string) string {
	value := strings.Replace(data, "\n", "", -1)
	return value
}
func removeSpace(data string) string {
	value := strings.Replace(data, " ", "", -1)
	return value
}
