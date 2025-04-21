package nrand

import (
	"io"
)

// ZeroReader is a globally available io.Reader that always overwrites zeros
var ZeroReader io.Reader

// Reader is a globally available io.Reader that performs noop
var Reader io.Reader

// Zero Reader container.
type Zero struct{}

// Noop Reader container.
type Noop struct{}

func init() {
	ZeroReader = &Zero{}
	Reader = &Noop{}
}

// Read for Zero
func (z Zero) Read(b []byte) (int, error) {
	l := len(b)
	copy(b, make([]byte, l))
	return l, nil
}

// Read for Noop
func (n Noop) Read(b []byte) (int, error) { return len(b), nil }

// ReadZeros take a byte slice and overwrites all bytes with zeros. It always returns len(b), nil.
func ReadZeros(b []byte) (int, error) {
	return io.ReadFull(ZeroReader, b)
}

// Read is a noop io.Reader. It always returns a len(b), nil.
func Read(b []byte) (int, error) {
	return io.ReadFull(Reader, b)
}

func genBytes(l int) []byte {
	b := make([]byte, l)
	Read(b)
	return b
}
