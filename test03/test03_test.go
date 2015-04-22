package test03

import (
	"testing"
)

type BigStruct struct {
	C01 int64
	C02 int64
	C03 int64
	C04 int64
	C05 int64
	C06 int64
	C07 int64
	C08 int64
	C09 int64
	C10 int64
}

func t1(a []BigStruct) {
	for i := 0; i < len(a); i++ {
		x, y := a[i].C01, a[i].C02
		//go func() {
		_ = x + y
		//}()
	}
}
func t2(a []BigStruct) {
	for i := 0; i < len(a); i++ {
		s := struct{ x, y int64 }{a[i].C01, a[i].C02}

		//go func() {
		_ = s.x + s.y
		//}()
	}
}

func Benchmark_01(b *testing.B) {
	b.StopTimer()
	var a = make([]BigStruct, 1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t1(a)
	}
}

func Benchmark_02(b *testing.B) {
	b.StopTimer()
	var a = make([]BigStruct, 1000)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t2(a)
	}
}
