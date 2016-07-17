#include "go_hello_module.h"
#include "redismodule.h"
#include <string.h>

/* ECHO <string> - Echo back a string sent from the client */
int EchoCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
  if (argc < 2) return RedisModule_WrongArity(ctx);
  RedisModule_AutoMemory(ctx);
  size_t len;
  char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));
  GoString *r = GoEcho(dst);
  RedisModuleString *rm_str = RedisModule_CreateString(ctx, (*r).p, (*r).n);
  return RedisModule_ReplyWithString(ctx, rm_str);
}

/* ECHO2 <string> - Echo back a string sent from the client */
int Echo2Command(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
  if (argc < 2) return RedisModule_WrongArity(ctx);
  RedisModule_AutoMemory(ctx);
  size_t len;
  char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));
  GoString *r = GoEcho2(dst, len);
  RedisModuleString *rm_str = RedisModule_CreateString(ctx, (*r).p, (*r).n);
  return RedisModule_ReplyWithString(ctx, rm_str);
}

/* ECHO3 <string> - Echo back a string sent from the client */
int Echo3Command(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
  if (argc < 2) return RedisModule_WrongArity(ctx);
  RedisModule_AutoMemory(ctx);
  size_t len;
  char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));
  char *r = GoEcho3(dst, len);
  size_t n = strlen(r);
  RedisModuleString *rm_str = RedisModule_CreateString(ctx, r, n);
  // free memory allocated in Golang stack
  free(r);
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
  if (RedisModule_CreateCommand(ctx, "hello.echo2", Echo2Command, "readonly", 1,1,1) == REDISMODULE_ERR) {
    return REDISMODULE_ERR;
  }
  if (RedisModule_CreateCommand(ctx, "hello.echo3", Echo3Command, "readonly", 1,1,1) == REDISMODULE_ERR) {
    return REDISMODULE_ERR;
  }
}
