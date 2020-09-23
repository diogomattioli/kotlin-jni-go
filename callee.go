package main

// #cgo CFLAGS: -I/usr/lib/jvm/java-14-openjdk-amd64/include
// #cgo CFLAGS: -I/usr/lib/jvm/java-14-openjdk-amd64/include/linux
// #include <jni.h>
import "C"

//export Java_Foo_add
func Java_Foo_add(env *C.JNIEnv, clazz C.jclass, a C.jint, b C.jint) C.jint {
	return a + b
}

//export Java_CallerKt_mul
func Java_CallerKt_mul(env *C.JNIEnv, clazz C.jclass, a C.jint, b C.jint) C.jint {
	return a * b
}

func main() {

}
