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

func ZeroCopySlice(p unsafe.Pointer, n int) *Slice {
	data := &c_slice_t{p, n}
	runtime.SetFinalizer(data, func(data *c_slice_t) {
		C.free(data.p)
	})
	s := &Slice{data: data}
	h := (*reflect.SliceHeader)((unsafe.Pointer(&s.Data)))
	h.Cap = n
	h.Len = n
	h.Data = uintptr(p)
	return s
}

//export GoEcho
// c->go: Convert C string to Go string
// go->c: return Go string pointer
func GoEcho(s *C.char) *string {
	gostr := (C.GoString(s) + " from golang1")
	return &gostr
}

//export GoEcho2
// c->go: Convert C array with explicit length to Go []byte using C.GoBytes
// go->c: return Go string pointer
func GoEcho2(s *C.char, length C.int) *string {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, " from golang2"...)
	gostr := string(slice)
	return &gostr
}

//export GoEcho3
// c->go: Convert C array with explicit length to Go []byte using C.GoBytes
// go->c: malloc the C buffer, and make a single copy into that buffer. It is
// important to keep in mind that the Go garbage collector will not interact
// with this data, and that if it is freed from the C side of things
func GoEcho3(s *C.char, length C.int) unsafe.Pointer {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, " from golang3"...)

	p := C.malloc(C.size_t(len(slice)))
	// free memory in c code
	// defer C.free(p)

	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], slice)

	return p
}

//export GoEcho4
// c->go: Convert C array with explicit length to Go []byte using C.GoBytes
// go->c: return unsafe.Pointer and length
func GoEcho4(s *C.char, length C.int) (unsafe.Pointer, int) {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, " from golang4"...)
	return unsafe.Pointer(&(slice[0])), len(slice)
}

//export GoEcho5
// error handler example
func GoEcho5(s *C.char) (*string, *string) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var gostr, err string
	if r.Intn(2) == 1 {
		gostr = ""
		err = "random system error"
	} else {
		gostr = (C.GoString(s) + " from golang5")
		err = ""
	}
	return &gostr, &err
}

//export GoEcho6
// c->go: using reflect to bind C memory to Go resource without memory copy.
// go->c: return unsafe.Pointer and length
func GoEcho6(s *C.char, length C.int) (unsafe.Pointer, int) {
	zslice := ZeroCopySlice(unsafe.Pointer(s), int(length))
	slice := append(zslice.Data, " from golang6"...)
	return unsafe.Pointer(&(slice[0])), len(slice)
}

func main() {}
