#! /usr/bin/bash

go run main.go test.sp

llc -O=0 --filetype=obj code.ll

clang code.o -o run

chmod +x run

rm code.*
