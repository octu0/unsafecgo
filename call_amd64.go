package unsafecgo

import (
	"unsafe"
)

// Call C func return will void:
//
//	void C.foo()
//
//go:noescape
func Call(unsafe.Pointer)

// Call C func return will int32:
//
//	int C.foo()
//
//go:noescape
func CallRC(unsafe.Pointer) int32

// Call C func return will pointer:
//
//	void* C.foo()
//
//go:noescape
func CallRP(unsafe.Pointer) unsafe.Pointer

// Call C func with context parameter and return will void:
//
//	void C.foo(void *ctx)
//
//go:noescape
func CallCtx(unsafe.Pointer, unsafe.Pointer)

// Call C func with context parameter and return will int32:
//
//	int C.foo(void *ctx)
//
//go:noescape
func CallCtxRC(unsafe.Pointer, unsafe.Pointer) int32

// Call C func with context parameter and return will pointer:
//
//	void* C.foo(void *ctx)
//
//go:noescape
func CallCtxRP(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
