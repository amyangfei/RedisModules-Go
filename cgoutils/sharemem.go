package cgoutils

// #include <stdlib.h>
import "C"

import (
	"reflect"
	"runtime"
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
