package utils

import (
	"log"
	"net/http"

	"github.com/avct/uasurfer"
)

// GetFile is use to download file
func GetFile(url string) *http.Response {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	return response
}

// IsChrome Check is user using chrome or not
func IsChrome(ua string) bool {
	browser := uasurfer.Parse(ua).Browser.Name.String()
	if browser == "BrowserChrome" {
		return true
	}
	return false
}
