#include "textflag.h"

// func Call(unsafe.Pointer)
TEXT ·Call(SB), NOSPLIT, $0
  MOVQ cfunc+0(FP), AX
  MOVQ SP, BX
  SUBQ $16, SP
  ANDQ $~15, SP
  CALL AX
  MOVQ BX, SP
  MOVQ AX, ret+8(FP)
  RET
