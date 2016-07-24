#!/bin/bash

[ -f redismodule.h ] || wget https://raw.githubusercontent.com/antirez/redis/unstable/src/redismodule.h
go build -tags gm -buildmode=c-shared -o go_rm_graphicsmagick.o rm_graphicsmagick.go
gcc -fPIC -std=gnu99 -c -o rm_graphicsmagick.o rm_graphicsmagick.c
ld -o go_rm_graphicsmagick.so rm_graphicsmagick.o $(pwd)/go_rm_graphicsmagick.o -shared -Bsymbolic -lc
