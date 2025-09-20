package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type ProxyWrap struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewProxy(rawURL string, timeout time.Duration) (*ProxyWrap, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	p := httputil.NewSingleHostReverseProxy(parsed)

	p.Transport = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ResponseHeaderTimeout: timeout,
	}

	return &ProxyWrap{
		target: parsed,
		proxy:  p,
	}, nil
}

func (pw *ProxyWrap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pw.proxy.ServeHTTP(w, r)
}
