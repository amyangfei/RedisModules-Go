package main

// #include <stdlib.h>
// #include "redismodule.h"
import "C"

import (
	"math/rand"
	"reflect"
	"runtime"
	"time"
	"unsafe"
)

type c_slice_t struct {
	p unsafe.Pointer
	n int
}

type Slice struct {
	Data []byte
	data *c_slice_t
}

func ZeroCopySlice(p unsafe.Pointer, length, capacity int, autofree bool) *Slice {
	data := &c_slice_t{p, length}
	if autofree {
		runtime.SetFinalizer(data, func(data *c_slice_t) {
			C.free(data.p)
		})
	}
	s := &Slice{data: data}
	h := (*reflect.SliceHeader)((unsafe.Pointer(&s.Data)))
	h.Cap = capacity
	h.Len = length
	h.Data = uintptr(p)
	return s
}

//export GoEcho1
// c->go: Convert C string to Go string
// go->c: Create C string via C.CString, return pointer and length of string
func GoEcho1(s *C.char) (*C.char, int) {
	gostr := (C.GoString(s) + " from golang1")
	return C.CString(gostr), len(gostr)
}

//export GoEcho2
// c->go: Convert C array with explicit length to Go []byte using C.GoBytes
// go->c: Create C string via C.CString, return pointer and length of string
func GoEcho2(s *C.char, length C.int) (*C.char, int) {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, " from golang2"...)
	return C.CString(string(slice)), len(slice)
}

//export GoEcho3
// c->go: Convert C array with explicit length to Go []byte using C.GoBytes
// go->c: malloc the C buffer, and make a single copy into that buffer. It is
// important to keep in mind that the Go garbage collector will not interact
// with this data, and that if it is freed from the C side of things
func GoEcho3(s *C.char, length C.int) (unsafe.Pointer, int) {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, " from golang3"...)

	p := C.malloc(C.size_t(len(slice)))
	// free memory in c code
	// defer C.free(p)

	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], slice)

	return p, len(slice)
}

//export GoEcho4
// c->go: using reflect to bind C memory to Go resource without memory copy.
// Fullfill c memory directly
func GoEcho4(s *C.char, length C.int) (unsafe.Pointer, int) {
	incr := " from golang4"
	cap := int(length) + len(incr)
	zslice := ZeroCopySlice(unsafe.Pointer(s), int(length), cap, false)
	copy(zslice.Data[int(length):cap], incr)
	return unsafe.Pointer(&zslice.Data[0]), cap
}

//export GoEcho5
// error handler example
func GoEcho5(s *C.char) (*C.char, int, *C.char, int) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	if r.Intn(2) == 1 {
		err := "random system error"
		return C.CString(""), 0, C.CString(err), len(err)
	} else {
		gostr := (C.GoString(s) + " from golang5")
		return C.CString(gostr), len(gostr), C.CString(""), 0
	}
}

//export GoEcho6
// c->go: using reflect to bind C memory to Go resource without memory copy.
// Fullfill c memory directly
func GoEcho6(s *C.char, length, capacity C.int) (unsafe.Pointer, int) {
	incr := " from golang6"
	zslice := ZeroCopySlice(unsafe.Pointer(s), int(capacity), int(capacity), false)
	copy(zslice.Data[int(length):], incr)
	return unsafe.Pointer(&zslice.Data[0]), int(length) + len(incr)
}

//export GoEcho7
// c->go: using reflect to bind C memory to Go resource without memory copy.
// go->c: malloc the C buffer, and make a single copy into that buffer.
func GoEcho7(s *C.char, length C.int) (unsafe.Pointer, int) {
	incr := " from golang7"
	// cap := int(length) + len(incr)
	zslice := ZeroCopySlice(unsafe.Pointer(s), int(length), int(length), false)
	zslice.Data = append(zslice.Data, incr...)

	p := C.malloc(C.size_t(len(zslice.Data)))
	// free memory in c code
	// defer C.free(p)
	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], zslice.Data)

	return p, len(zslice.Data)
}

func main() {}
