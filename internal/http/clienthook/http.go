package clienthook

import "net/http"

type transport struct {
	transport http.RoundTripper
	hook      func(*http.Request)
}

func NewTransport(rt http.RoundTripper, hook func(*http.Request)) *transport {
	if rt == nil {
		rt = http.DefaultTransport
	}

	if hook == nil {
		hook = func(*http.Request) {}
	}

	return &transport{
		transport: rt,
		hook:      hook,
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.hook(req)
	return t.transport.RoundTrip(req)
}
