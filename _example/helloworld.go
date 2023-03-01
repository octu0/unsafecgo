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
	"time"
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

func benchmark_cgo(N int) {
	t := time.Now()
	for i := 0; i < N; i += 1 {
		C.hello()
	}
	fmt.Println("cgo(avg)", time.Since(t)/time.Duration(N))
}

func benchmark_unsafe(N int) {
	t := time.Now()
	p := unsafe.Pointer(reflect.ValueOf(C.hello).Pointer())

	for i := 0; i < N; i += 1 {
		unsafecgo.Call(p)
	}

	runtime.KeepAlive(p)
	fmt.Println("unsafe(avg)", time.Since(t)/time.Duration(N))
}

func benchmark() {
	N := 5
	benchmark_cgo(N)
	benchmark_unsafe(N)
}
