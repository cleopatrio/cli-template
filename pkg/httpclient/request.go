package httpclient

import (
	"bytes"
	"io"
)

type RequestBody struct {
	Data  []byte
	state bytes.Buffer
}

func (rb *RequestBody) Close() error { return io.NopCloser(&rb.state).Close() }

func (rb *RequestBody) Read(buffer []byte) (int, error) { return rb.state.Read(buffer) }

// RequestBody implements io.ReadCloser
var _ io.ReadCloser = (*RequestBody)(nil)
