#include "go_password.h"
#include "redismodule.h"
#include <string.h>


static int cmd_password_set(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    RedisModuleCallReply *reply;
    size_t len;
    char *hash;

    if (argc != 3) {
        RedisModule_WrongArity(ctx);
        return REDISMODULE_OK;
    }

    char *password = RedisModule_Strdup(RedisModule_StringPtrLen(argv[2], &len));
    struct GoDoCrypt_return r = GoDoCrypt(password, len);

    if ((*r.r1).n != 0) {
        RedisModule_ReplyWithError(ctx, (*r.r1).p);
        return REDISMODULE_ERR;
    }
    reply = RedisModule_Call(ctx, "SET", "sc!", argv[1], (*r.r0).p);
    RedisModule_ReplyWithCallReply(ctx, reply);

    return REDISMODULE_OK;
}


static int cmd_password_hset(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    return REDISMODULE_OK;
}


static int cmd_password_check(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    return REDISMODULE_OK;
}


static int cmd_password_hcheck(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    return REDISMODULE_OK;
}


int RedisModule_OnLoad(RedisModuleCtx *ctx) {
    if (RedisModule_Init(ctx, "go_password", 1, REDISMODULE_APIVER_1) == REDISMODULE_ERR)
        return REDISMODULE_ERR;

    if (RedisModule_CreateCommand(ctx, "go_password.set", cmd_password_set,
                "no-monitor fast", 1, 1, 1) == REDISMODULE_ERR)
        return REDISMODULE_ERR;
    if (RedisModule_CreateCommand(ctx, "go_password.hset", cmd_password_hset,
                "no-monitor fast", 1, 1, 1) == REDISMODULE_ERR)
        return REDISMODULE_ERR;
    if (RedisModule_CreateCommand(ctx, "go_password.check", cmd_password_check,
                "no-monitor fast", 1, 1, 1) == REDISMODULE_ERR)
        return REDISMODULE_ERR;
    if (RedisModule_CreateCommand(ctx, "go_password.hcheck", cmd_password_hcheck,
                "no-monitor fast", 1, 1, 1) == REDISMODULE_ERR)
        return REDISMODULE_ERR;

    return REDISMODULE_OK;
}
