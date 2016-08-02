package main

/*
#include "redismodule.h"

inline RedisModuleString *RedisModule_CreateString_Wrap(RedisModuleCtx *ctx, char *ptr, size_t len) {
    void *getapifuncptr = ((void**)ctx)[0];
    RedisModule_GetApi = (int (*)(const char *, void *)) (unsigned long)getapifuncptr;
    RedisModule_GetApi("RedisModule_CreateString", (void **)&RedisModule_CreateString);

    RedisModuleString *rms = RedisModule_CreateString(ctx, ptr, len);
    return rms;
}
*/
import "C"

//export GoEcho
func GoEcho(ctx *C.RedisModuleCtx, s *C.char) *C.RedisModuleString {
	gostr := (C.GoString(s) + " from golang version bridge")
	return C.RedisModule_CreateString_Wrap(ctx, C.CString(gostr), C.size_t(len(gostr)))
}

func main() {}
