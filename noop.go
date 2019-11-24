package noop

// Void does nothing.
func Void() {}

// Error returns nil.
func Error() error { return nil }

// IntErr returns 0, nil.
func IntErr() (int, error) { return 0, nil }

// Int64Err returns 0, nil.
func Int64Err() (int64, error) { return 0, nil }

// ByteErr returns 0, nil.
func ByteErr() (byte, error) { return 0, nil }

//  Noop is a struct with general purpose methods
type Noop struct{}

// Error returns ""
func (n Noop) Error() string { return "" }

// String returns "".
func (n Noop) String() string { return "" }
