package nnet

import (
	"net"
	"time"
)

// Noop is a struct that has methods that conform to all `net` interfaces
type Noop struct{}

// Network returns "".
func (n Noop) Network() string { return "" }

// Read returns 0, nil.
func (n Noop) Read([]byte) (int, error) { return 0, nil }

// Write returns 0, nil.
func (n Noop) Write([]byte) (int, error) { return 0, nil }

// Close returns nil.
func (n Noop) Close() error { return nil }

// String returns "".
func (n Noop) String() string { return "" }

// LocalAddr returns nil.
func (n Noop) LocalAddr() net.Addr { return nil }

// RemoteAddr returns nil.
func (n Noop) RemoteAddr() net.Addr { return nil }

// SetDeadline returns nil.
func (n Noop) SetDeadline(time.Time) error { return nil }

// SetReadDeadline returns nil.
func (n Noop) SetReadDeadline(time.Time) error { return nil }

// SetWriteDeadline returns nil.
func (n Noop) SetWriteDeadline(time.Time) error { return nil }

// Error returns "".
func (n Noop) Error() string { return "" }

// Timeout returns false.
func (n Noop) Timeout() bool { return false }

// Temporary returns false
func (n Noop) Temporary() bool { return false }

// Accept retuns nil, nil
func (n Noop) Accept() (net.Conn, error) { return nil, nil }

// Addr returns Noop{}
func (n Noop) Addr() net.Addr { return nil }

// ReadFrom returns 0, nil, nil.
func (n Noop) ReadFrom([]byte) (int, net.Addr, error) { return 0, nil, nil }

// WriteTo returns 0, nil.
func (n Noop) WriteTo([]byte, net.Addr) (int, error) { return 0, nil }
