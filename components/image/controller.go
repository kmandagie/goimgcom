package image

import (
	"bytes"
	"image/jpeg"
	"image/png"
	"log"

	"github.com/chai2010/webp"
	"github.com/h2non/filetype"
	"github.com/kmandagie/goimgcom/components/utils"
	"github.com/labstack/echo"
)

// Errordata error message
type Errordata struct {
	Code    string `json:"code"`
	Message string `message:"email"`
}

// Optimize will optimize automatically will return image value
func Optimize(c echo.Context) error {
	url := c.ParamValues()[0]
	userAgent := c.Request().Header.Get("User-Agent")
	image := utils.GetFile(url)
	buff := utils.ReaderToBuffer(image.Body)
	if !isImage(buff) {
		message := &Errordata{
			Code:    "400",
			Message: "Filetype not valid",
		}
		return c.JSON(400, message)
	}
	if utils.IsChrome(userAgent) {
		webp, mime := convertToWebp(buff)
		return c.Stream(200, mime, bytes.NewReader(webp))
	}
	mime := utils.GetMimeType(buff)
	return c.Stream(200, mime, bytes.NewReader(buff))
}

func convertToWebp(buff []byte) ([]byte, string) {
	if isJpeg(buff) {
		m, err := jpeg.Decode(bytes.NewReader(buff))
		if err != nil {
			log.Fatal(err)
		}
		data, err := webp.EncodeRGB(m, 75.0)
		if err != nil {
			log.Fatal(err)
		}
		return data, "image/webp"
	}
	if isPng(buff) {
		m, err := png.Decode(bytes.NewReader(buff))
		if err != nil {
			log.Fatal(err)
		}
		data, err := webp.EncodeRGB(m, 75.0)
		if err != nil {
			log.Fatal(err)
		}
		return data, "image/webp"
	}

	return buff, utils.GetMimeType(buff)
}

func isImage(buff []byte) bool {
	if filetype.IsImage(buff) {
		return true
	}
	return false
}

func isJpeg(buff []byte) bool {
	mime := utils.GetMimeType(buff)
	if mime == "image/jpeg" {
		return true
	}
	return false
}

func isPng(buff []byte) bool {
	mime := utils.GetMimeType(buff)
	if mime == "image/png" {
		return true
	}
	return false
}
