package benchmark

import (
	"testing"
)

func BenchmarkUnsafecgo(b *testing.B) {
	b.Run("cgo/malloc_free", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			cgo_malloc_free()
		}
	})
	b.Run("unsafecgo/malloc_free", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			unsafecgo_malloc_free()
		}
	})

	b.Run("cgo/calc", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			cgo_calc()
		}
	})
	b.Run("unsafecgo/calc", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			unsafecgo_calc()
		}
	})

	b.Run("cgo/nop_call", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			cgo_nop_call()
		}
	})
	b.Run("unsafecgo/nop_call", func(tb *testing.B) {
		for i := 0; i < tb.N; i += 1 {
			unsafecgo_nop_call()
		}
	})
}
