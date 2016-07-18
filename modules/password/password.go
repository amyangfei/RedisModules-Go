package main

// #include "redismodule.h"
import "C"

import (
	"golang.org/x/crypto/bcrypt"
	"unsafe"
)

//export GoDoCrypt
func GoDoCrypt(password *C.char, length C.int) (*string, *string) {
	var retHash, retErr string

	passwdSlice := C.GoBytes(unsafe.Pointer(password), length)

	// Hashing the password with cost of 5
	hashed, err := bcrypt.GenerateFromPassword(passwdSlice, 5)
	if err != nil {
		retHash, retErr = "", err.Error()
	} else {
		retHash, retErr = string(hashed), ""
	}

	return &retHash, &retErr
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
