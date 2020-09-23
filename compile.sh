#!/bin/bash
rm -rf build
mkdir build
kotlinc-jvm -include-runtime -d build/foo.jar *.kt
javah -d build/ -cp build/foo.jar Foo CallerKt
go build -buildmode=c-shared -o build/libfoo.so *.go