#include "textflag.h"

// func Available() bool
TEXT ·Available(SB),NOSPLIT,$0
  MOVL $1, AX        // true
  MOVB AX, ret+0(FP)
  RET

