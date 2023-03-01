package unsafecgo

import (
	"unsafe"
)

// call: void C.foo()
//
//go:noescape
func Call(unsafe.Pointer)

// call: int C.foo()
//
//go:noescape
func CallRC(unsafe.Pointer) int32

// call: int C.foo(void *ctx)
//
//go:noescape
func CallCtxRC(unsafe.Pointer, unsafe.Pointer) int32
