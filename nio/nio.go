package nio

import "io"

// Noop is a struct used as a container for all `io` methods to satisfy all `io` interfaces.
type Noop struct{}

// Close returns nil.
func (n Noop) Close() error { return nil }

// Read returns 0, nil.
func (n Noop) Read([]byte) (int, error) { return 0, nil }

// ReadAt returns 0, nil.
func (n Noop) ReadAt([]byte, int64) (int, error) { return 0, nil }

// ReadByte returns 0, nil.
func (n Noop) ReadByte() (byte, error) { return 0, nil }

// ReadFrom returns 0, nil.
func (n Noop) ReadFrom(io.Reader) (int64, error) { return 0, nil }

// ReadRune returns 0, 0, nil.
func (n Noop) ReadRune() (rune, int, error) { return 0, 0, nil }

// Seek returns 0, nil.
func (n Noop) Seek(int64, int) (int64, error) { return 0, nil }

// UnreadByte returns nil.
func (n Noop) UnreadByte() error { return nil }

// UnreadRune returns nil.
func (n Noop) UnreadRune() error { return nil }

// Write returns 0, nil.
func (n Noop) Write([]byte) (int, error) { return 0, nil }

// WriteAt return 0, nil.
func (n Noop) WriteAt([]byte, int64) (int, error) { return 0, nil }

// WriteString returns 0, nil.
func (n Noop) WriteString(string) (int, error) { return 0, nil }

// WriteTo returns 0, nil.
func (n Noop) WriteTo(io.Writer) (int64, error) { return 0, nil }
