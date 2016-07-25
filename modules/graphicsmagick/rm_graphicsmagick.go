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
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	err = mw.ThumbnailImage(uint(width), uint(height))
	if err != nil {
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	retImg, retErr = mw.GetImagesBlob(), ""

	return unsafe.Pointer(&(retImg[0])), len(retImg), &retErr
}

//export GoImgRotate
func GoImgRotate(buf *C.char, length C.int, degrees C.double) (unsafe.Pointer, int, *string) {
	var retImg []byte
	var retErr string

	data := C.GoBytes(unsafe.Pointer(buf), length)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(data)
	if err != nil {
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	background := imagick.NewPixelWand()
	err = mw.RotateImage(background, float64(degrees))
	if err != nil {
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	retImg, retErr = mw.GetImagesBlob(), ""

	return unsafe.Pointer(&(retImg[0])), len(retImg), &retErr
}

//export GoImgBlur
func GoImgBlur(buf *C.char, length C.int, radius, sigma C.double) (unsafe.Pointer, int, *string) {
	var retImg []byte
	var retErr string

	data := C.GoBytes(unsafe.Pointer(buf), length)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(data)
	if err != nil {
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	err = mw.BlurImage(float64(radius), float64(sigma))
	if err != nil {
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	retImg, retErr = mw.GetImagesBlob(), ""

	return unsafe.Pointer(&(retImg[0])), len(retImg), &retErr
}

//export GoImgSwirl
func GoImgSwirl(buf *C.char, length C.int, degrees C.double) (unsafe.Pointer, int, *string) {
	var retImg []byte
	var retErr string

	data := C.GoBytes(unsafe.Pointer(buf), length)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(data)
	if err != nil {
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	err = mw.SwirlImage(float64(degrees))
	if err != nil {
		retImg, retErr = []byte("m"), err.Error()
		return unsafe.Pointer(&(retImg[0])), 0, &retErr
	}

	retImg, retErr = mw.GetImagesBlob(), ""

	return unsafe.Pointer(&(retImg[0])), len(retImg), &retErr
}

func main() {}
