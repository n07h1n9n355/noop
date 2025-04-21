package nrand

import (
	"testing"
)

func BenchmarkRead(b *testing.B) {
	Read(make([]byte, 500000000))
}

func BenchmarkReadZeros(b *testing.B) {
	ReadZeros(make([]byte, 500000000))
}
