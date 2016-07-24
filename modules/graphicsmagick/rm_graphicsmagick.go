package main

// #include "redismodule.h"
import "C"

import (
	"bytes"
	"github.com/rainycape/magick"
	"unsafe"
)

//export GoGetImgType
func GoGetImgType(buf *C.char, length C.int) (*string, *string) {
	var imgType, retErr string
	data := C.GoBytes(unsafe.Pointer(buf), length)
	img, err := magick.DecodeData(data)
	if err != nil {
		imgType, retErr = "", err.Error()
	} else {
		imgType, retErr = img.Format(), ""
	}
	return &imgType, &retErr
}

//export GoImgThumbnail
func GoImgThumbnail(buf *C.char, length C.int, width, height C.longlong) (*string, *string) {
	var retImg, retErr string

	data := C.GoBytes(unsafe.Pointer(buf), length)

	img, err := magick.DecodeData(data)
	if err != nil {
		retImg, retErr = "", err.Error()
		return &retImg, &retErr
	}

	thumbnail, err := img.Thumbnail(int(width), int(height))
	if err != nil {
		retImg, retErr = "", err.Error()
		return &retImg, &retErr
	}

	var wbuf bytes.Buffer
	err = thumbnail.Encode(&wbuf, nil)
	if err != nil {
		retImg, retErr = "", err.Error()
		return &retImg, &retErr
	}

	retImg, retErr = wbuf.String(), ""

	return &retImg, &retErr
}

func main() {}
