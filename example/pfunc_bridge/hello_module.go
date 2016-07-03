package main

// #include "redismodule.h"
/*
typedef RedisModuleString *(*redis_func) (RedisModuleCtx *ctx, char *ptr, size_t len);

inline RedisModuleString *redis_bridge_func(redis_func f, RedisModuleCtx *ctx, char *ptr, size_t len)
{
	return f(ctx, ptr, len);
}
*/
import "C"

//export GoEcho
func GoEcho(ctx *C.RedisModuleCtx, s *C.char) *C.RedisModuleString {
	gostr := (C.GoString(s) + " from golang version 2")
	f := C.redis_func(C.RedisModule_CreateString)
	return C.redis_bridge_func(f, ctx, C.CString(gostr), C.size_t(len(gostr)))
}

func main() {}
