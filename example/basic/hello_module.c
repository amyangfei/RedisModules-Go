#include "go_hello_module.h"
#include "redismodule.h"
#include <string.h>

/* ECHO <string> - Echo back a string sent from the client */
int EchoCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc < 2) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);
    size_t len;
    char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));

    struct GoEcho_return r = GoEcho(dst);
    RedisModuleString *rm_str = RedisModule_CreateString(ctx, r.r0, r.r1);
    free(r.r0);

    RedisModule_ReplyWithString(ctx, rm_str);
    return REDISMODULE_OK;
}

/* ECHO2 <string> - Echo back a string sent from the client */
int Echo2Command(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc < 2) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);
    size_t len;
    char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));

    struct GoEcho2_return r = GoEcho2(dst, len);
    RedisModuleString *rm_str = RedisModule_CreateString(ctx, r.r0, r.r1);
    free(r.r0);

    RedisModule_ReplyWithString(ctx, rm_str);
    return REDISMODULE_OK;
}

/* ECHO3 <string> - Echo back a string sent from the client */
int Echo3Command(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc < 2) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);
    size_t len;
    char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));
    struct GoEcho3_return r = GoEcho3(dst, len);
    RedisModuleString *rm_str = RedisModule_CreateString(ctx, r.r0, r.r1);
    // free memory allocated in Golang stack
    free(r.r0);
    RedisModule_ReplyWithString(ctx, rm_str);
    return REDISMODULE_OK;
}

/* ECHO4 <string> - Echo back a string sent from the client */
int Echo4Command(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc < 2) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);
    size_t len;
    char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));
    /* RedisModule_Realloc(dst, capacity); */
    struct GoEcho4_return r = GoEcho4(dst, len);
    RedisModuleString *rm_str = RedisModule_CreateString(ctx, r.r0, r.r1);
    RedisModule_ReplyWithString(ctx, rm_str);
    return REDISMODULE_OK;
}

/* ECHO5 <string> - Echo back a string sent from the client */
int Echo5Command(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc < 2) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);
    size_t len;
    char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));
    struct GoEcho5_return r = GoEcho5(dst);
    if (r.r3 != 0 ) {
        RedisModule_ReplyWithError(ctx, r.r2);
        free(r.r0);
        free(r.r2);
        return REDISMODULE_ERR;
    } else {
        RedisModuleString *rm_str = RedisModule_CreateString(ctx, r.r0, r.r1);
        RedisModule_ReplyWithString(ctx, rm_str);
        free(r.r0);
        free(r.r2);
        return REDISMODULE_OK;
    }
}

/* ECHO6 <string> - Echo back a string sent from the client */
int Echo6Command(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc < 2) return RedisModule_WrongArity(ctx);
    RedisModule_AutoMemory(ctx);
    size_t len;
    char *dst = RedisModule_Strdup(RedisModule_StringPtrLen(argv[1], &len));
    size_t capacity = len + 20;
    RedisModule_Realloc(dst, capacity);
    struct GoEcho6_return r = GoEcho6(dst, len, capacity);
    RedisModuleString *rm_str = RedisModule_CreateString(ctx, r.r0, r.r1);
    RedisModule_ReplyWithString(ctx, rm_str);
    return REDISMODULE_OK;
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
    if (RedisModule_CreateCommand(ctx, "hello.echo4", Echo4Command, "readonly", 1,1,1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }
    if (RedisModule_CreateCommand(ctx, "hello.echo5", Echo5Command, "readonly", 1,1,1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }
    if (RedisModule_CreateCommand(ctx, "hello.echo6", Echo6Command, "readonly", 1,1,1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }
}
