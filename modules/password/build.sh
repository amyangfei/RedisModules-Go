#!/bin/bash

[ -f redismodule.h ] || wget https://raw.githubusercontent.com/antirez/redis/unstable/src/redismodule.h
go build -buildmode=c-shared -o go_password.o password.go
gcc -fPIC -std=gnu99 -c -o password.o password.c
ld -o go_password.so password.o $(pwd)/go_password.o -shared -Bsymbolic -lc
