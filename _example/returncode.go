package main

/*
#include <stdio.h>
#include <stdlib.h>

int return_code() {
  return -100;
}

typedef struct param_t {
  unsigned char *in;
  int size;
  unsigned char *out;
} param_t;

param_t *create_param_t() {
  return (param_t*)malloc(sizeof(param_t));
}

int convert(void *ctx) {
  param_t *p = (param_t*)(ctx);

  for(int i = 0; i < p->size; i += 1) {
    p->out[i] = p->in[i] + 100;
  }

  return 0;
}
*/
import "C"

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"

	"github.com/octu0/unsafecgo"
)

var (
	ptrReturnCodeFunc  = unsafe.Pointer(reflect.ValueOf(C.return_code).Pointer())
	ptrCreateParamFunc = unsafe.Pointer(reflect.ValueOf(C.create_param_t).Pointer())
	ptrConvertFunc     = unsafe.Pointer(reflect.ValueOf(C.convert).Pointer())
	ptrFreeFunc        = unsafe.Pointer(reflect.ValueOf(C.free).Pointer())
)

func cgo_return_code() {
	ret := int32(C.return_code())
	fmt.Println("cgo:return_code", ret)
}

func unsafecgo_return_code() {
	ret := unsafecgo.CallRC(ptrReturnCodeFunc)
	fmt.Println("unsafecgo:return_code", ret)
}

func cgo_convert() {
	t := time.Now()
	defer fmt.Println("cgo elapsed", time.Since(t))

	input := []byte{1, 2, 3, 4, 5}
	output := make([]byte, len(input))
	param := C.create_param_t()
	param.in = (*C.uchar)(unsafe.Pointer(&input[0]))
	param.size = C.int(len(input))
	param.out = (*C.uchar)(unsafe.Pointer(&output[0]))

	ret := int32(C.convert(unsafe.Pointer(param)))
	fmt.Println("cgo:convert/rc", ret)
	fmt.Println("cgo:convert/out", output)
	C.free(unsafe.Pointer(param))
}

func unsafecgo_convert() {
	t := time.Now()
	defer fmt.Println("unsafecgo elapsed", time.Since(t))

	input := []byte{1, 2, 3, 4, 5}
	output := make([]byte, len(input))
	param := (*C.param_t)(unsafecgo.CallRP(ptrCreateParamFunc))
	param.in = (*C.uchar)(unsafe.Pointer(&input[0]))
	param.size = C.int(len(input))
	param.out = (*C.uchar)(unsafe.Pointer(&output[0]))

	ret := unsafecgo.CallCtxRC(ptrConvertFunc, unsafe.Pointer(param))
	fmt.Println("cgo:convert/rc", ret)
	fmt.Println("cgo:convert/out", output)
	unsafecgo.CallCtx(ptrFreeFunc, unsafe.Pointer(param))
}

func main() {
	cgo_return_code()
	unsafecgo_return_code()

	cgo_convert()
	unsafecgo_convert()
}
