package context

import (
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	honojson "github.com/evex-dev/hono.go/src/json"
)

func (c *Context) ParseBody() ([]byte, error) {
	return io.ReadAll(c.ReadBody())
}

func (c *Context) ParseText() (string, error) {
	body, err := c.ParseBody()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (c *Context) ParseJson(target any) error {
	body, err := c.ParseBody()
	if err != nil {
		return err
	}
	return honojson.ParseJSON(string(body), &target)
}

func (c *Context) GetParam(key string) string {
	value, ok := c.Params[key]
	if !ok {
		return fmt.Sprintf("param '%s' not found", key)
	}
	return value
}

func (c *Context) Method() string {
	return c.Req.Method
}

func (c *Context) Query() url.Values {
	return c.Req.URL.Query()
}

func (c *Context) Host() string {
	return c.Req.Host
}

func (c *Context) URL() *url.URL {
	return c.Req.URL
}

func (c *Context) Proto() string {
	return c.Req.Proto
}

func (c *Context) ProtoMajor() int {
	return c.Req.ProtoMajor
}

func (c *Context) ProtoMinor() int {
	return c.Req.ProtoMajor
}

func (c *Context) ReadBody() io.ReadCloser {
	return c.Req.Body
}

func (c *Context) GetBody() func() (io.ReadCloser, error) {
	return c.Req.GetBody
}

func (c *Context) ContentLength() int64 {
	return c.Req.ContentLength
}

func (c *Context) TransferEncoding() []string {
	return c.Req.TransferEncoding
}

func (c *Context) Close() bool {
	return c.Req.Close
}

func (c *Context) Form() url.Values {
	return c.Req.Form
}

func (c *Context) PostForm() url.Values {
	return c.Req.PostForm
}

func (c *Context) MultipartForm() *multipart.Form {
	return c.Req.MultipartForm
}

func (c *Context) RemoteAddr() string {
	return c.Req.RemoteAddr
}

func (c *Context) ReqURI() string {
	return c.Req.RequestURI
}

func (c *Context) TLS() *tls.ConnectionState {
	return c.Req.TLS
}

func (c *Context) Cancel() <-chan struct{} {
	return c.Req.Cancel
}

func (c *Context) Response() *http.Response {
	return c.Req.Response
}
