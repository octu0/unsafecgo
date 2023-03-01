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
```

# License

MIT, see LICENSE file for details.
