## RedisModules

This repo introduces how to write redis loadable modules with golang.

There are also some redis loadable modules written in C.

## Contents

### example

[example/basic](./example/basic) simple echo module written in go

[example/pfunc_bridge](./example/pfunc_bridge) simple echo module written in go, use function bridge to call API defined in `redismoule.h` in golang code.

[example/benchmark](./example/benchmark) simple benchmark for all modules in example/basic

### redis modules written with golang

[modules/password](./modules/password) fork of [pasword](http://redismodules.com/modules/password/), written in go

[modules/graphicsmagick](./modules/graphicsmagick) fork of [graphicsmagick](http://redismodules.com/modules/graphicsmagick-2/), written in go

### redis modules written with C

[c-modules/cmd_router](./c-modules/cmd_router) a command router  between redis client and redis server
