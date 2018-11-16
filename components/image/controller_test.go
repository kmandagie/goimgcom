package image

import (
	"testing"

	"github.com/kmandagie/goimgcom/components/utils"
)

func TestIsJpegTrue(t *testing.T) {
	image := utils.GetFile("https://c5.staticflickr.com/9/8738/16720197924_3a0e2e9657_o.jpg")
	buff := utils.ReaderToBuffer(image.Body)
	isjpeg := isJpeg(buff)
	if !isjpeg {
		t.Errorf("Image is  must Jpeg, got: %t, want: %t.", isjpeg, true)
	}

}

// Icon Test Taken from https://icons8.com
func TestIsJpegFalse(t *testing.T) {
	image := utils.GetFile("https://img.icons8.com/metro/50/000000/facebook.png")
	buff := utils.ReaderToBuffer(image.Body)
	isjpeg := isJpeg(buff)
	if isjpeg {
		t.Errorf("Image is must not Jpeg, got: %t, want: %t.", isjpeg, false)
	}

}

// Icon Test Taken from https://icons8.com
func TestIsPNGTrue(t *testing.T) {
	image := utils.GetFile("https://cdn.detik.net.id/detik2/images/logodetikcom.png")
	buff := utils.ReaderToBuffer(image.Body)
	ispng := isPng(buff)
	if !ispng {
		t.Errorf("Image is must png, got: %t, want: %t.", ispng, true)
	}

}
func TestIsPNGFalse(t *testing.T) {
	image := utils.GetFile("https://c5.staticflickr.com/9/8738/16720197924_3a0e2e9657_o.jpg")
	buff := utils.ReaderToBuffer(image.Body)
	ispng := isPng(buff)
	if ispng {
		t.Errorf("Image is must not png, got: %t, want: %t.", ispng, false)
	}

}

func TestConvertJpegTowebp(t *testing.T) {
	image := utils.GetFile("https://c5.staticflickr.com/9/8738/16720197924_3a0e2e9657_o.jpg")
	imagedata := utils.ReaderToBuffer(image.Body)
	buff, _ := convertToWebp(imagedata)
	mime := utils.GetMimeType(buff)
	if mime != "image/webp" {
		t.Errorf("Image jpeg must convert to webp , got: %s, want: %s.", mime, "image/webp")
	}
}

// Icon Test Taken from https://icons8.com
func TestConvertPngTowebp(t *testing.T) {
	image := utils.GetFile("https://img.icons8.com/metro/50/000000/facebook.png")
	imagedata := utils.ReaderToBuffer(image.Body)
	buff, _ := convertToWebp(imagedata)
	mime := utils.GetMimeType(buff)
	if mime != "image/webp" {
		t.Errorf("Image png must convert to webp , got: %s, want: %s.", mime, "image/webp")
	}
}

// Giphy GIF for testing only
func TestNotConvertTowebp(t *testing.T) {
	image := utils.GetFile("https://media.giphy.com/media/3oD3YveOJWdwIAfZ5e/giphy.gif")
	imagedata := utils.ReaderToBuffer(image.Body)
	buff, _ := convertToWebp(imagedata)
	mime := utils.GetMimeType(buff)
	if mime == "image/webp" {
		t.Errorf("Image other than jpeg and png must not convert to webp ")
	}
}

func TestIsImageTrue(t *testing.T) {
	image := utils.GetFile("https://c5.staticflickr.com/9/8738/16720197924_3a0e2e9657_o.jpg")
	buff := utils.ReaderToBuffer(image.Body)
	isimage := isImage(buff)
	if !isimage {
		t.Errorf("file is must image, got: %t, want: %t.", isimage, true)
	}

}

// Currently not supported svg
func TestIsImageFalse(t *testing.T) {
	image := utils.GetFile("https://assets-cdn.github.com/pinned-octocat.svg")
	buff := utils.ReaderToBuffer(image.Body)
	isimage := isImage(buff)
	if isimage {
		t.Errorf("file is must not image, got: %t, want: %t.", isimage, false)
	}

}
