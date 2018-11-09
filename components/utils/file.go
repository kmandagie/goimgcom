package utils

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/h2non/filetype"
)

// ReaderToBuffer Convert Reader IO to buffer so we can use the data multiple times
func ReaderToBuffer(r io.Reader) []byte {
	buff, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return buff
}

// GetMimeType get file type of extensions return string
func GetMimeType(buff []byte) string {
	kind, err := filetype.Match(buff)
	if err != nil {
		log.Fatal(err)
	}
	return kind.MIME.Value
}
