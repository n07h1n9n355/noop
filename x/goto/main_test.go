package main

import (
	"testing"
)

func benchmarkRecFibN(i uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		recFib(i)
	}
}

func benchmarkFibN(i uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(i)
	}
}

func benchmarkGibN(i uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		gotoFib(i)
	}
}

func benchmarkIterFibN(i uint64, b *testing.B) {
	for n := 0; n < b.N; n++ {
		impFib(i)
	}
}

// // func BenchmarkRecFib50(b *testing.B) { benchmarkRecFibN(100, b) }

func BenchmarkFib1(b *testing.B)    { benchmarkFibN(1, b) }
func BenchmarkFib100(b *testing.B)  { benchmarkFibN(100, b) }
func BenchmarkFib1600(b *testing.B) { benchmarkFibN(1600, b) }

// func BenchmarkFib100(b *testing.B) { benchmarkFibN(100, b) }

func BenchmarkIterFib1(b *testing.B)    { benchmarkIterFibN(1, b) }
func BenchmarkIterFib20(b *testing.B)   { benchmarkIterFibN(20, b) }
func BenchmarkIterFib100(b *testing.B)  { benchmarkIterFibN(100, b) }
func BenchmarkIterFib200(b *testing.B)  { benchmarkIterFibN(200, b) }
func BenchmarkIterFib400(b *testing.B)  { benchmarkIterFibN(400, b) }
func BenchmarkIterFib800(b *testing.B)  { benchmarkIterFibN(800, b) }
func BenchmarkIterFib1600(b *testing.B) { benchmarkIterFibN(1600, b) }

func BenchmarkGib1(b *testing.B)    { benchmarkGibN(1, b) }
func BenchmarkGib20(b *testing.B)   { benchmarkGibN(20, b) }
func BenchmarkGib100(b *testing.B)  { benchmarkGibN(100, b) }
func BenchmarkGib200(b *testing.B)  { benchmarkGibN(200, b) }
func BenchmarkGib400(b *testing.B)  { benchmarkGibN(400, b) }
func BenchmarkGib800(b *testing.B)  { benchmarkGibN(800, b) }
func BenchmarkGib1600(b *testing.B) { benchmarkGibN(1600, b) }

// func BenchmarkRecFib1(b *testing.B)  { benchmarkRecFibN(1, b) }
// func BenchmarkRecFib20(b *testing.B) { benchmarkRecFibN(20, b) }
