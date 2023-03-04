//go:build !amd64
// +build !amd64

package unsafecgo

import (
	"unsafe"
)

func Call(unsafe.Pointer) {
	panic("unavailable Call")
}

func CallRC(unsafe.Pointer) int32 {
	panic("unavailable CallRC")
}

func CallRP(unsafe.Pointer) unsafe.Pointer {
	panic("unavailable CallRP")
}

func CallCtx(unsafe.Pointer, unsafe.Pointer) {
	panic("unavailable CallCtx")
}

func CallCtxRC(unsafe.Pointer, unsafe.Pointer) int32 {
	panic("unavailable CallCtxRC")
}

func CallCtxRP(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer {
	panic("unavailable CallCtxRC")
}
