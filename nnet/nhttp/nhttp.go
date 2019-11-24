package nhttp

import (
	"bufio"
	"net"
	"net/http"
	"net/url"
	"os"
)

// NoopCloseNotifier conforms to the responseWriter CloserNotifier interface
type NoopCloseNotifier struct {
	closeNotifyCh chan bool
}

// NewCloseNotifier returns a pointer to a Noop
func NewCloseNotifier() *NoopCloseNotifier {
	return &NoopCloseNotifier{
		closeNotifyCh: make(chan bool, 1),
	}
}

// CloseNotify returns  <-chan bool
func (n *NoopCloseNotifier) CloseNotify() <-chan bool { return n.closeNotifyCh }

// Noop is used for `http` interfaces
type Noop struct{}

// SetCookies does nothing.
func (n Noop) SetCookies(*url.URL, []*http.Cookie) {}

// Cookies returns an empty slice of pointers to http.Cookie
func (n Noop) Cookies(*url.URL) []*http.Cookie { return []*http.Cookie{&http.Cookie{}} }

// Readdir returns []os.FileInfo{nil}, nil
func (n Noop) Readdir(int) ([]os.FileInfo, error) { return []os.FileInfo{nil}, nil }

// Stat returns nil, nil
func (n Noop) Stat() (os.FileInfo, error) { return nil, nil }

// Close returns nil.
func (n Noop) Close() error { return nil }

// Read returns 0, nil.
func (n Noop) Read([]byte) (int, error) { return 0, nil }

// Seek returns 0, nil.
func (n Noop) Seek(int64, int) (int64, error) { return 0, nil }

// Open returns nil, nil.
func (n Noop) Open(string) (http.File, error) { return nil, nil }

// Flush does nothing.
func (n Noop) Flush() {}

// ServeHTTP does nothing.
func (n Noop) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

// Hijack returns nil, nil, nil
func (n Noop) Hijack() (net.Conn, *bufio.ReadWriter, error) { return nil, nil, nil }

// Push returns nil
func (n Noop) Push(string, *http.PushOptions) error { return nil }

// Header returns an empty map[string][]string
func (n Noop) Header() http.Header { return make(map[string][]string) }

// Write returns 0, nil.
func (n Noop) Write([]byte) (int, error) { return 0, nil }

// WriteHeader does nothing.
func (n Noop) WriteHeader(int) {}

// RoundTrip returns nil, nil
func (n Noop) RoundTrip(*http.Request) (*http.Response, error) { return nil, nil }
