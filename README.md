# `unsafecgo`

[![MIT License](https://img.shields.io/github/license/octu0/unsafecgo)](https://github.com/octu0/unsafecgo/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/octu0/unsafecgo?status.svg)](https://godoc.org/github.com/octu0/unsafecgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/octu0/unsafecgo)](https://goreportcard.com/report/github.com/octu0/unsafecgo)
[![Releases](https://img.shields.io/github/v/release/octu0/unsafecgo)](https://github.com/octu0/unsafecgo/releases)

`unsafecgo` provides Cgo calls via assembly trampoline. Inspired by [rustgo](https://blog.filippo.io/rustgo/).

## Installation

```bash
go get github.com/octu0/unsafecgo
```

## Example

```go
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

func cgo_helloworld() {
	fmt.Println("cgo")

	C.hello()
}

func unsafecgo_helloworld() {
	fmt.Println("unsafecgo.Call")

	p := unsafe.Pointer(reflect.ValueOf(C.hello).Pointer())
	unsafecgo.Call(p) // call C.hello()
	runtime.KeepAlive(p)
}

func main() {
	cgo_helloworld()
	unsafecgo_helloworld()
}
```

## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/octu0/unsafecgo/benchmark
cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz
BenchmarkUnsafecgo
BenchmarkUnsafecgo/cgo/malloc_free
BenchmarkUnsafecgo/cgo/malloc_free-4              971646              1236 ns/op            1024 B/op          1 allocs/op
BenchmarkUnsafecgo/unsafecgo/malloc_free
BenchmarkUnsafecgo/unsafecgo/malloc_free-4       1340142               910.4 ns/op             0 B/op          0 allocs/op
BenchmarkUnsafecgo/cgo/calc
BenchmarkUnsafecgo/cgo/calc-4                    3667410               311.5 ns/op             0 B/op          0 allocs/op
BenchmarkUnsafecgo/unsafecgo/calc
BenchmarkUnsafecgo/unsafecgo/calc-4             16357767                74.31 ns/op            0 B/op          0 allocs/op
BenchmarkUnsafecgo/cgo/nop_call
BenchmarkUnsafecgo/cgo/nop_call-4               17337052                65.60 ns/op            0 B/op          0 allocs/op
BenchmarkUnsafecgo/unsafecgo/nop_call
BenchmarkUnsafecgo/unsafecgo/nop_call-4         369709251                3.211 ns/op           0 B/op          0 allocs/op
PASS
```

# License

MIT, see LICENSE file for details.
