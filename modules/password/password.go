package main

// #include "redismodule.h"
import "C"

import (
	"github.com/amyangfei/RedisModules-Go/cgoutils"
	"golang.org/x/crypto/bcrypt"
	"unsafe"
)

//export GoDoCrypt
func GoDoCrypt(password *C.char, length C.int) (*C.char, int, *C.char, int) {
	passwdSlice := cgoutils.ZeroCopySlice(
		unsafe.Pointer(password), int(length), int(length), false)

	// Hashing the password with cost of 5
	hashed, err := bcrypt.GenerateFromPassword(passwdSlice.Data, 5)
	if err != nil {
		errstr := err.Error()
		return C.CString(""), 0, C.CString(errstr), len(errstr)
	} else {
		return C.CString(string(hashed)), len(hashed), C.CString(""), 0
	}
}

//export GoDoValidate
func GoDoValidate(hash, password *C.char, hashLen, passLen C.int) int {
	hashSlice := C.GoBytes(unsafe.Pointer(hash), hashLen)
	passSlice := C.GoBytes(unsafe.Pointer(password), passLen)

	err := bcrypt.CompareHashAndPassword(hashSlice, passSlice)
	if err != nil {
		return 0
	} else {
		return 1
	}
}

func main() {}
