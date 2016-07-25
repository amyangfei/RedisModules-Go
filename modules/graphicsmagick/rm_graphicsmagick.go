package main

// #include "redismodule.h"
import "C"

import (
	"gopkg.in/gographics/imagick.v2/imagick"
	"unsafe"
)

//export GoGetImgType
func GoGetImgType(buf *C.char, length C.int) (*string, *string) {
	var imgType, retErr string
	data := C.GoBytes(unsafe.Pointer(buf), length)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(data)
	if err != nil {
		imgType, retErr = "", err.Error()
	} else {
		imgType, retErr = mw.GetImageFormat(), ""
	}

	return &imgType, &retErr
}

//export GoImgThumbnail
func GoImgThumbnail(buf *C.char, length C.int, width, height C.longlong) (unsafe.Pointer, int, *string) {
	var retImg []byte
	var retErr string

	data := C.GoBytes(unsafe.Pointer(buf), length)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(data)
	if err != nil {
		retImg, retErr = []byte(""), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	err = mw.ThumbnailImage(uint(width), uint(height))
	if err != nil {
		retImg, retErr = []byte(""), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	retImg, retErr = mw.GetImagesBlob(), ""

	return unsafe.Pointer(&(retImg[0])), len(retImg), &retErr
}

func main() {}
