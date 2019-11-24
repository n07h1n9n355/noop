package nos

import (
	"os"
	"time"
)

// Noop is an empty struct with methods that conform to all interfaces in the `os` package
type Noop struct{}

// Name returns "".
func (n Noop) Name() string { return "" }

// Size returns 0.
func (n Noop) Size() int64 { return 0 }

// Mode returns 0
func (n Noop) Mode() os.FileMode { return 0 }

// ModTime returns time.Time{}.
func (n Noop) ModTime() time.Time { return time.Time{} }

// IsDir returns false.
func (n Noop) IsDir() bool { return false }

// Sys returns nil.
func (n Noop) Sys() interface{} { return nil }

// String returns "".
func (n Noop) String() string { return "" }

// Signal returns nothing.
func (n Noop) Signal() {}
