package wrap

import "net/http"

// Wrapper can wrap a http.Handler with another one
type Wrapper interface {
	// Wrap wraps the next `http.Handler` of the stack and returns the previous `http.Handler`
	// If `next.ServeHTTP` is not called, the next `http.Handler` won't be used
	Wrap(next http.Handler) (previous http.Handler)
}

// noop is a http.Handler that does nothing
var noop = http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})

// New returns a http.Handler that wraps the given wrappers.
// When it serves the request the first given wrapper
// serves the request and may let the second wrapper (its "next" wrapper) serve.
// The second wrapper may let the third wrapper serve and so on.
// The last wrapper has as "next" wrapper the not exported noop handler that does nothing.
func New(wrapper ...Wrapper) (h http.Handler) {
	h = noop
	for i := len(wrapper) - 1; i >= 0; i-- {
		h = wrapper[i].Wrap(h)
	}
	return
}
