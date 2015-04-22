package test02

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

func t1() int64 {
	var r BigStruct
	r.C01 = 0
	r.C02 = 0
	r.C03 = 0
	r.C04 = 0
	r.C05 = 0
	r.C06 = 0
	r.C07 = 0
	r.C08 = 0
	r.C09 = 0
	r.C10 = 0
	return r.C01 + r.C02 + r.C03 + r.C04 + r.C05 + r.C06 + r.C07 + r.C08 + r.C09 + r.C10
}

func t2() int64 {
	var c01, c02, c03, c04, c05, c06, c07, c08, c09, c10 int64
	c01 = 0
	c02 = 0
	c03 = 0
	c04 = 0
	c05 = 0
	c06 = 0
	c07 = 0
	c08 = 0
	c09 = 0
	c10 = 0
	return c01 + c02 + c03 + c04 + c05 + c06 + c07 + c08 + c09 + c10

}
func t3(r BigStruct) int64 {
	r.C01 = 0
	r.C02 = 0
	r.C03 = 0
	r.C04 = 0
	r.C05 = 0
	r.C06 = 0
	r.C07 = 0
	r.C08 = 0
	r.C09 = 0
	r.C10 = 0
	return r.C01 + r.C02 + r.C03 + r.C04 + r.C05 + r.C06 + r.C07 + r.C08 + r.C09 + r.C10
}

func Benchmark_01(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t1()
	}
}

func Benchmark_02(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t2()
	}
}
func Benchmark_03(b *testing.B) {
	b.StopTimer()
	var r BigStruct
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		t3(r)
	}
}
