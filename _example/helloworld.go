package main

/*
#include <stdio.h>
void hello() {
  fprintf(stderr, "hello world\n");
}
*/
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/octu0/unsafecgo"
)

func helloworld() {
	p := unsafe.Pointer(reflect.ValueOf(C.hello).Pointer())

	unsafecgo.Call(p)

	runtime.KeepAlive(p)
}

func main() {
	fmt.Println("cgo")
	C.hello()

	fmt.Println("unsafecgo.Call")
	helloworld()
}
