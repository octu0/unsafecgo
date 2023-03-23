#include "textflag.h"

// func Call(unsafe.Pointer)
TEXT ·Call(SB), NOSPLIT, $0
  MOVQ cfunc+0(FP), AX
  MOVQ SP, BX
  SUBQ $16, SP
  ANDQ $~15, SP
  CALL AX
  MOVQ BX, SP
  MOVQ AX, ret+0(FP)
  RET

// func CallRC(unsafe.Pointer) int32
TEXT ·CallRC(SB), NOSPLIT, $0
  MOVQ cfunc+0(FP), AX
  MOVQ SP, BX
  SUBQ $16, SP
  ANDQ $~15, SP
  CALL AX
  MOVQ BX, SP
  MOVQ AX, ret+8(FP)
  RET

// func CallRP(unsafe.Pointer) unsafe.Pointer
TEXT ·CallRP(SB), NOSPLIT, $0
  MOVQ cfunc+0(FP), AX
  MOVQ SP, BX
  SUBQ $16, SP
  ANDQ $~15, SP
  CALL AX
  MOVQ BX, SP
  MOVQ AX, ret+8(FP)
  RET

// func CallCtx(unsafe.Pointer, unsafe.Pointer)
TEXT ·CallCtx(SB), 0, $2048-16
  MOVQ cfunc+0(FP), AX
  MOVQ arg0+8(FP), DI
  MOVQ SP, BX
  ADDQ $2048, SP
  ANDQ $~15, SP
  CALL AX
  MOVQ BX, SP
  MOVQ AX, ret+16(FP)
  RET

// func CallCtxRC(unsafe.Pointer, unsafe.Pointer) int32
TEXT ·CallCtxRC(SB), 0, $2048-16
  MOVQ cfunc+0(FP), AX
  MOVQ arg0+8(FP), DI
  MOVQ SP, BX
  ADDQ $2048, SP
  ANDQ $~15, SP
  CALL AX
  MOVQ BX, SP
  MOVQ AX, ret+16(FP)
  RET

// func CallCtxRP(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
TEXT ·CallCtxRP(SB), 0, $2048-16
  MOVQ cfunc+0(FP), AX
  MOVQ arg0+8(FP), DI
  MOVQ SP, BX
  ADDQ $2048, SP
  ANDQ $~15, SP
  CALL AX
  MOVQ BX, SP
  MOVQ AX, ret+16(FP)
  RET
