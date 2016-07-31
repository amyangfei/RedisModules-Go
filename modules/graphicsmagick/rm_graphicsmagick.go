package main

// #include "redismodule.h"
import "C"

import (
	"github.com/amyangfei/RedisModules-Go/cgoutils"
	"gopkg.in/gographics/imagick.v2/imagick"
	"unsafe"
)

func formatResult(result interface{}, err error) (*C.char, int, *C.char, int) {
	if err != nil {
		errstr := err.Error()
		return C.CString(""), 0, C.CString(errstr), len(errstr)
	}
	switch v := result.(type) {
	default:
		errstr := "unexpected result type"
		return C.CString(""), 0, C.CString(errstr), len(errstr)
	case []byte:
		return C.CString(string(v)), len(v), C.CString(""), 0
	case string:
		return C.CString(v), len(v), C.CString(""), 0
	}
}

//export GoGetImgType
func GoGetImgType(buf *C.char, length C.int) (*C.char, int, *C.char, int) {
	raw := cgoutils.ZeroCopySlice(unsafe.Pointer(buf), int(length), int(length), false)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(raw.Data)
	if err != nil {
		return formatResult(nil, err)
	} else {
		return formatResult(mw.GetImageFormat(), nil)
	}
}

//export GoImgThumbnail
func GoImgThumbnail(buf *C.char, length C.int, width, height C.longlong) (*C.char, int, *C.char, int) {
	raw := cgoutils.ZeroCopySlice(unsafe.Pointer(buf), int(length), int(length), false)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(raw.Data)
	if err != nil {
		return formatResult(nil, err)
	}

	err = mw.ThumbnailImage(uint(width), uint(height))
	if err != nil {
		return formatResult(nil, err)
	}

	return formatResult(mw.GetImageBlob(), nil)
}

//export GoImgRotate
func GoImgRotate(buf *C.char, length C.int, degrees C.double) (*C.char, int, *C.char, int) {
	raw := cgoutils.ZeroCopySlice(unsafe.Pointer(buf), int(length), int(length), false)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(raw.Data)
	if err != nil {
		return formatResult(nil, err)
	}

	background := imagick.NewPixelWand()
	err = mw.RotateImage(background, float64(degrees))
	if err != nil {
		return formatResult(nil, err)
	}

	return formatResult(mw.GetImageBlob(), nil)
}

//export GoImgBlur
func GoImgBlur(buf *C.char, length C.int, radius, sigma C.double) (*C.char, int, *C.char, int) {
	raw := cgoutils.ZeroCopySlice(unsafe.Pointer(buf), int(length), int(length), false)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(raw.Data)
	if err != nil {
		return formatResult(nil, err)
	}

	err = mw.BlurImage(float64(radius), float64(sigma))
	if err != nil {
		return formatResult(nil, err)
	}

	return formatResult(mw.GetImageBlob(), nil)
}

//export GoImgSwirl
func GoImgSwirl(buf *C.char, length C.int, degrees C.double) (*C.char, int, *C.char, int) {
	raw := cgoutils.ZeroCopySlice(unsafe.Pointer(buf), int(length), int(length), false)

	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()

	err := mw.ReadImageBlob(raw.Data)
	if err != nil {
		return formatResult(nil, err)
	}

	err = mw.SwirlImage(float64(degrees))
	if err != nil {
		return formatResult(nil, err)
	}

	return formatResult(mw.GetImageBlob(), nil)
}

func main() {}
