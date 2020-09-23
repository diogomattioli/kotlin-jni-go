# kotlin-jni-go

**Compiling Kotlin**

```kotlinc-jvm -include-runtime -d build/foo.jar *.kt```

**Generate the header files**

```javah -d build/ -cp build/foo.jar Foo CallerKt```

**Compiling Go**

```go build -buildmode=c-shared -o build/libfoo.so *.go```

**Running...**

```java -Djava.library.path=build/ -cp build/foo.jar CallerKt```
