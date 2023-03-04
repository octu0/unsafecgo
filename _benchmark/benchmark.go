package benchmark

/*
#include <stdio.h>
#include <stdlib.h>

void nop_func() {
  // noop
}

typedef struct param_calc_t {
  int a;
  int b;
  int out;
} param_calc_t;

param_calc_t *create_param_calc() {
  return (param_calc_t*)malloc(sizeof(param_calc_t));
}

int calc(void *ctx) {
  param_calc_t *param = (param_calc_t *)ctx;

  param->out = param->a * param->b;

  return 0;
}

int update_bytes(unsigned char *data, int size) {
  for(int i = 0; i < size; i += 1) {
    data[i] = 123;
  }
  return 0;
}

typedef struct wrap_malloc_param_t {
  int size;
} wrap_malloc_param_t;

wrap_malloc_param_t *create_malloc_param() {
  return (wrap_malloc_param_t*)malloc(sizeof(wrap_malloc_param_t));
}

void *wrap_malloc(void *ctx) {
  wrap_malloc_param_t *param = (wrap_malloc_param_t*)ctx;
  return malloc(param->size);
}

typedef struct wrap_param_t {
  void *data;
  int size;
} wrap_param_t;

wrap_param_t *create_param() {
  return (wrap_param_t*)malloc(sizeof(wrap_param_t));
}

int wrap_update_bytes(void *ctx) {
  wrap_param_t *param = (wrap_param_t *)ctx;
  return update_bytes((unsigned char *)param->data, param->size);
}
*/
import "C"

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/octu0/cgobytepool"
	"github.com/octu0/unsafecgo"
)

func cgo_nop_call() {
	C.nop_func()
}

func cgo_malloc_free() {
	data := C.malloc(1024)
	_ = C.update_bytes((*C.uchar)(data), C.int(1024))

	b := C.GoBytes(data, C.int(1024))
	for i := 0; i < 1024; i += 1 {
		if b[i] != 123 {
			panic("wrong value")
		}
	}

	C.free(data)

	runtime.KeepAlive(data)
}

func cgo_calc() {
	param := C.create_param_calc()
	param.a = C.int(1234)
	param.b = C.int(3)
	param.out = C.int(0)

	_ = C.calc(unsafe.Pointer(param))

	if param.out != C.int(3702) {
		panic("wrong value")
	}

	C.free(unsafe.Pointer(param))

	runtime.KeepAlive(param)
}

var (
	nopFunc = unsafe.Pointer(reflect.ValueOf(C.nop_func).Pointer())
)

func unsafecgo_nop_call() {
	unsafecgo.Call(nopFunc)

	runtime.KeepAlive(nopFunc)
}

var (
	createCalcParamFunc = unsafe.Pointer(reflect.ValueOf(C.create_param_calc).Pointer())
	calcFunc            = unsafe.Pointer(reflect.ValueOf(C.calc).Pointer())
	freeFunc            = unsafe.Pointer(reflect.ValueOf(C.free).Pointer())
)

func unsafecgo_calc() {
	paramPtr := unsafecgo.CallRP(createCalcParamFunc)
	param := (*C.param_calc_t)(paramPtr)
	param.a = 1234
	param.b = 3
	param.out = C.int(0)

	_ = unsafecgo.CallCtxRC(calcFunc, paramPtr)

	if param.out != C.int(3702) {
		panic("wrong value")
	}

	unsafecgo.CallCtx(freeFunc, paramPtr)

	runtime.KeepAlive(paramPtr)
}

var (
	wrapMallocFunc        = unsafe.Pointer(reflect.ValueOf(C.wrap_malloc).Pointer())
	createMallocParamFunc = unsafe.Pointer(reflect.ValueOf(C.create_malloc_param).Pointer())
	createParamFunc       = unsafe.Pointer(reflect.ValueOf(C.create_param).Pointer())
	wrapUpdateBytesFunc   = unsafe.Pointer(reflect.ValueOf(C.wrap_update_bytes).Pointer())
)

func unsafecgo_malloc_free() {
	// C.create_malloc_param
	mallocParamPtr := unsafecgo.CallRP(createMallocParamFunc)
	mallocParam := (*C.wrap_malloc_param_t)(mallocParamPtr)
	mallocParam.size = C.int(1024)

	// C.wrap_malloc
	dataPtr := unsafecgo.CallCtxRP(wrapMallocFunc, mallocParamPtr)

	// C.create_param
	paramPtr := unsafecgo.CallRP(createParamFunc)
	param := (*C.wrap_param_t)(paramPtr)
	param.data = dataPtr
	param.size = mallocParam.size

	// C.wrap_update_bytes
	_ = unsafecgo.CallCtxRC(wrapUpdateBytesFunc, paramPtr)

	b := cgobytepool.GoBytes(dataPtr, 1024)
	for i := 0; i < 1024; i += 1 {
		if b[i] != 123 {
			panic("wrong value")
		}
	}

	unsafecgo.CallCtx(freeFunc, paramPtr)
	unsafecgo.CallCtx(freeFunc, dataPtr)
	unsafecgo.CallCtx(freeFunc, mallocParamPtr)

	runtime.KeepAlive(mallocParamPtr)
	runtime.KeepAlive(dataPtr)
	runtime.KeepAlive(paramPtr)
}
