package main

// #include <stdlib.h>
// #include "redismodule.h"
import "C"

import (
	"math/rand"
	"time"
	"unsafe"
)

//export GoEcho
func GoEcho(s *C.char) *string {
	gostr := (C.GoString(s) + " from golang")
	return &gostr
}

//export GoEcho2
func GoEcho2(s *C.char, length C.int) *string {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, []byte(" from golang2")...)
	gostr := string(slice)
	return &gostr
}

//export GoEcho3
func GoEcho3(s *C.char, length C.int) unsafe.Pointer {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, []byte(" from golang3")...)

	p := C.malloc(C.size_t(len(slice)))
	// free memory in c code
	// defer C.free(p)

	cBuf := (*[1 << 30]byte)(p)
	copy(cBuf[:], slice)

	return p
}

//export GoEcho4
func GoEcho4(s *C.char, length C.int) (unsafe.Pointer, int) {
	slice := C.GoBytes(unsafe.Pointer(s), length)
	slice = append(slice, []byte(" from golang4")...)
	return unsafe.Pointer(&(slice[0])), len(slice)
}

//export GoEcho5
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

func main() {}
