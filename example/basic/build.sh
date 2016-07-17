#!/bin/bash

[ -f redismodule.h ] || wget https://raw.githubusercontent.com/antirez/redis/unstable/src/redismodule.h
go build -buildmode=c-shared -o go_hello_module.o hello_module.go
gcc -fPIC -std=gnu99 -c -o hello_module.o hello_module.c
ld -o hello_module.so hello_module.o $(pwd)/go_hello_module.o -shared -Bsymbolic -lc
