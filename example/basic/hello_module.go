package main

// #include "redismodule.h"
import "C"

//export GoEcho
func GoEcho(s *C.char) *string {
	gostr := (C.GoString(s) + " from golang")
	return &gostr
}

func main() {}
