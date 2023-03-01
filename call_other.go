//go:build !amd64
// +build !amd64

package unsafecgo

import (
	"runtime"
	"unsafe"
)

//go:linkname asmcgocall runtime.asmcgocall
//go:noescape
func asmcgocall(unsafe.Pointer, unsafe.Pointer) int32

//go:linkname entersyscall runtime.entersyscall
//go:noescape
func entersyscall()

//go:linkname exitsyscall runtime.exitsyscall
//go:noescape
func exitsyscall()

func call(fn unsafe.Pointer, arg unsafe.Pointer) int32 {
	entersyscall()
	ret := asmcgocall(fn, arg)
	exitsyscall()

	return ret
}

func Call(fn unsafe.Pointer) {
	call(fn, nil)
	runtime.KeepAplive(fn)
}

func CallRC(fn unsafe.Pointer) int32 {
	ret := call(fn, nil)
	runtime.KeepAplive(fn)

	return ret
}

func CallCtxRC(fn unsafe.Pointer, arg unsafe.Pointer) int32 {
	ret := call(fn, arg)
	runtime.KeepAplive(fn)

	return ret
}
