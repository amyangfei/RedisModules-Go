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
    RedisModuleCallReply *reply;
    char *password;
    size_t len;

    if (argc != 4) {
        RedisModule_WrongArity(ctx);
        return REDISMODULE_OK;
    }

    password = RedisModule_Strdup(RedisModule_StringPtrLen(argv[3], &len));
    struct GoDoCrypt_return r = GoDoCrypt(password, len);

    if ((*r.r1).n != 0) {
        RedisModule_ReplyWithError(ctx, (*r.r1).p);
        return REDISMODULE_ERR;
    }
    reply = RedisModule_Call(ctx, "HSET", "ssc!", argv[1], argv[2], (*r.r0).p);
    RedisModule_ReplyWithCallReply(ctx, reply);

    return REDISMODULE_OK;
}


static void validate(RedisModuleCtx *ctx, RedisModuleCallReply *stored_hash, RedisModuleString *password)
{
    char *stored_hash_str;
    size_t hash_len;
    char *pass;
    size_t pass_len;

    if (RedisModule_CallReplyType(stored_hash) == REDISMODULE_REPLY_NULL) {
        RedisModule_ReplyWithLongLong(ctx, 0);
        return;
    }

    if (RedisModule_CallReplyType(stored_hash) != REDISMODULE_REPLY_STRING) {
        RedisModule_ReplyWithError(ctx, "WRONGTYPE Operation against a key holding the wrong kind of value");
        return;
    }

    stored_hash_str = RedisModule_Strdup(RedisModule_CallReplyStringPtr(stored_hash, &hash_len));
    pass = RedisModule_Strdup(RedisModule_StringPtrLen(password, &pass_len));

    GoInt compare = GoDoValidate(stored_hash_str, pass, hash_len, pass_len);
    if (compare == 1) {
        RedisModule_ReplyWithLongLong(ctx, 1);
    } else {
        RedisModule_ReplyWithLongLong(ctx, 0);
    }
}

static int cmd_password_check(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    RedisModuleCallReply *stored_hash;

    if (argc != 3) {
        RedisModule_WrongArity(ctx);
        return REDISMODULE_OK;
    }

    stored_hash = RedisModule_Call(ctx, "GET", "s", argv[1]);
    validate(ctx, stored_hash, argv[2]);

    return REDISMODULE_OK;
}


static int cmd_password_hcheck(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    RedisModuleCallReply *reply;

    if (argc != 4) {
        RedisModule_WrongArity(ctx);
    return REDISMODULE_OK;
    }

    reply = RedisModule_Call(ctx, "HGET", "ss", argv[1], argv[2]);
    validate(ctx, reply, argv[3]);

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
