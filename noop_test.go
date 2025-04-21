package noop

import (
	"testing"
)

func BenchmarkVoid(b *testing.B) {
	Void()
}
