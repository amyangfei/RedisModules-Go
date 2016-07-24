#include "go_rm_graphicsmagick.h"
#include "redismodule.h"

#include <string.h>


static int UpdateTransformedImage(RedisModuleCtx *ctx, RedisModuleKey *img_key,
                                  const char *transformed_img, size_t img_len) {
    int res = REDISMODULE_ERR;

    // Resize the key to contain the transformed image buffer
    if (RedisModule_StringTruncate(img_key, img_len) == REDISMODULE_ERR) {
        RedisModule_ReplyWithError(ctx, "Error resizing string key");
        return res;
    }

    // Get write access to the resized key
    char *buf = RedisModule_StringDMA(img_key, &img_len, REDISMODULE_WRITE);
    if (!buf) {
        RedisModule_ReplyWithError(ctx, REDISMODULE_ERRORMSG_WRONGTYPE);
        return res;
    }

    // Write transformed blob to output
    memcpy(buf, transformed_img, img_len);

    // Setup success reply
    RedisModule_ReplyWithSimpleString(ctx, "OK");
    res = REDISMODULE_OK;

    return res;
}


int GMRotateCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc != 3) {
        return RedisModule_WrongArity(ctx);
    }

    // init auto memory for created strings
    RedisModule_AutoMemory(ctx);

    // Validate input is valid float
    double degrees;
    if (RedisModule_StringToDouble(argv[2], &degrees) == REDISMODULE_ERR) {
        RedisModule_ReplyWithError(ctx, "Invalid arguments");
        return REDISMODULE_ERR;
    }

    // TODO: rotate image

    return REDISMODULE_OK;
}

int GMSwirlCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc != 3) {
        return RedisModule_WrongArity(ctx);
    }

    // init auto memory for created strings
    RedisModule_AutoMemory(ctx);

    // Validate input is valid float
    double degrees;
    if (RedisModule_StringToDouble(argv[2], &degrees) == REDISMODULE_ERR) {
        RedisModule_ReplyWithError(ctx, "Invalid arguments");
        return REDISMODULE_ERR;
    }

    // TODO: swirl image

    return REDISMODULE_OK;
}

int GMBlurCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc != 4) {
        return RedisModule_WrongArity(ctx);
    }

    // init auto memory for created strings
    RedisModule_AutoMemory(ctx);

    // Validate inputs are valid floats
    double radius;
    if (RedisModule_StringToDouble(argv[2], &radius) == REDISMODULE_ERR) {
        RedisModule_ReplyWithError(ctx, "Invalid arguments");
        return REDISMODULE_ERR;
    }
    double sigma;
    if (RedisModule_StringToDouble(argv[3], &sigma) == REDISMODULE_ERR) {
        RedisModule_ReplyWithError(ctx, "Invalid arguments");
        return REDISMODULE_ERR;
    }

    // TODO: blur image

    return REDISMODULE_OK;
}

int GMThumbnailCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    if (argc != 4) {
        return RedisModule_WrongArity(ctx);
    }

    // init auto memory for created strings
    RedisModule_AutoMemory(ctx);

    // open the key
    RedisModuleKey *key = RedisModule_OpenKey(ctx, argv[1], REDISMODULE_READ | REDISMODULE_WRITE);
    // If key doesn't exist then return immediately
    if (RedisModule_KeyType(key) == REDISMODULE_KEYTYPE_EMPTY) {
        RedisModule_ReplyWithError(ctx, "empty key");
        return REDISMODULE_OK;
    }
    // Validate key is a string
    if (RedisModule_KeyType(key) != REDISMODULE_KEYTYPE_STRING) {
        RedisModule_ReplyWithError(ctx, REDISMODULE_ERRORMSG_WRONGTYPE);
        return REDISMODULE_ERR;
    }

    // Validate inputs are valid unsigned ints
    long long width;
    if (RedisModule_StringToLongLong(argv[2], &width) == REDISMODULE_ERR || width <= 0) {
        RedisModule_ReplyWithError(ctx, "Invalid arguments");
        return REDISMODULE_ERR;
    }
    long long height;
    if (RedisModule_StringToLongLong(argv[3], &height) == REDISMODULE_ERR || height <= 0) {
        RedisModule_ReplyWithError(ctx, "Invalid arguments");
        return REDISMODULE_ERR;
    }

    // Get access to the image
    size_t key_len;
    char *buf = RedisModule_StringDMA(key, &key_len, REDISMODULE_READ);
    if (!buf) {
        RedisModule_ReplyWithError(ctx, REDISMODULE_ERRORMSG_WRONGTYPE);
        return REDISMODULE_ERR;
    }

    struct GoImgThumbnail_return r = GoImgThumbnail(buf, key_len, width, height);
    if ((*r.r1).n != 0) {
        RedisModule_ReplyWithError(ctx, (*r.r1).p);
        return REDISMODULE_ERR;
    }
    UpdateTransformedImage(ctx, key, (*r.r0).p, (*r.r0).n);

    return REDISMODULE_OK;
}

int GMGetTypeCommand(RedisModuleCtx *ctx, RedisModuleString **argv, int argc) {
    int res = REDISMODULE_ERR;

    if (argc != 2) {
        return RedisModule_WrongArity(ctx);
    }

    // init auto memory for created strings
    RedisModule_AutoMemory(ctx);

    // open the key
    RedisModuleKey *key = RedisModule_OpenKey(ctx, argv[1], REDISMODULE_READ);

    // If key doesn't exist then return immediately
    if (RedisModule_KeyType(key) == REDISMODULE_KEYTYPE_EMPTY) {
        RedisModule_ReplyWithError(ctx, "empty key");
        return REDISMODULE_OK;
    }

    // Validate key is a string
    if (RedisModule_KeyType(key) != REDISMODULE_KEYTYPE_STRING) {
        RedisModule_ReplyWithError(ctx, REDISMODULE_ERRORMSG_WRONGTYPE);
        return REDISMODULE_ERR;
    }

    // Get access to the image
    size_t key_len;
    char *buf = RedisModule_StringDMA(key, &key_len, REDISMODULE_READ);
    if (!buf) {
        RedisModule_ReplyWithError(ctx, REDISMODULE_ERRORMSG_WRONGTYPE);
        return REDISMODULE_ERR;
    }

    struct GoGetImgType_return r = GoGetImgType(buf, key_len);

    if ((*r.r1).n != 0) {
        RedisModule_ReplyWithError(ctx, (*r.r1).p);
        return REDISMODULE_ERR;
    }
    RedisModule_ReplyWithSimpleString(ctx, (*r.r0).p);

    return  REDISMODULE_OK;
}

int RedisModule_OnLoad(RedisModuleCtx *ctx) {

    // Register the module itself
    if (RedisModule_Init(ctx, "go_graphicsmagick", 1, REDISMODULE_APIVER_1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    // Register the commands
    if (RedisModule_CreateCommand(ctx, "go_graphicsmagick.rotate", GMRotateCommand, "write", 1, 1, 1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    if (RedisModule_CreateCommand(ctx, "go_graphicsmagick.swirl", GMSwirlCommand, "write", 1, 1, 1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    if (RedisModule_CreateCommand(ctx, "go_graphicsmagick.blur", GMBlurCommand, "write", 1, 1, 1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    if (RedisModule_CreateCommand(ctx, "go_graphicsmagick.thumbnail", GMThumbnailCommand, "write", 1, 1, 1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    if (RedisModule_CreateCommand(ctx, "go_graphicsmagick.type", GMGetTypeCommand, "readonly", 1, 1, 1) == REDISMODULE_ERR) {
        return REDISMODULE_ERR;
    }

    return REDISMODULE_OK;
}
