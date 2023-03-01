//go:build !amd64
// +build !amd64

//
package unsafecgo

import (
	"unsafe"
)

func Available() bool {
	return false
}
