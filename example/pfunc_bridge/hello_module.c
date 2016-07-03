#include "go_hello_module.h"
#include "redismodule.h"
#include <string.h>

/* ECHO <string> - Echo back a string sent from the client */
int EchoCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
  if (argc < 2) return RedisModule_WrongArity(ctx);
  RedisModule_AutoMemory(ctx);
  size_t len;
  const char *str = RedisModule_StringPtrLen(argv[1], &len);
  char *dst = strndup(str, len);
  RedisModuleString *rm_str = GoEcho(ctx, dst);
  return RedisModule_ReplyWithString(ctx, rm_str);
}

/* Registering the module */
int RedisModule_OnLoad(RedisModuleCtx *ctx) {
  if (RedisModule_Init(ctx, "hello", 1, REDISMODULE_APIVER_1) == REDISMODULE_ERR) {
    return REDISMODULE_ERR;
  }
  if (RedisModule_CreateCommand(ctx, "hello.echo", EchoCommand, "readonly", 1,1,1) == REDISMODULE_ERR) {
    return REDISMODULE_ERR;
  }
}
