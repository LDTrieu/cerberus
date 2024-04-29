package xtransport

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/ldtrieu/cerberus/package/cache"
)

type cacheMethodGetTransport struct {
	cacheClient cache.ICache
	transport   http.RoundTripper
}

func NewGetCache(cacheClient cache.ICache) http.RoundTripper {
	return &cacheMethodGetTransport{
		cacheClient: cacheClient,
		transport:   http.DefaultTransport,
	}
}

func cacheKey(r *http.Request) string {
	return fmt.Sprintf("xtransport:%s", r.URL.String())
}

func (c *cacheMethodGetTransport) get(r *http.Request) ([]byte, error) {
	val, err := c.cacheClient.Get(context.Background(), cacheKey(r))
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c *cacheMethodGetTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if r.Method == http.MethodGet {
		if val, err := c.get(r); err == nil {
			buf := bytes.NewBuffer(val)
			return http.ReadResponse(bufio.NewReader(buf), r)
		}
	}

	resp, err := c.transport.RoundTrip(r)
	if err != nil {
		return nil, err
	}

	// Get the body of the response so we can save it in the cache for the next request.
	buf, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		c.cacheClient.Set(context.Background(), cacheKey(r), buf)
	}

	return resp, nil
}
