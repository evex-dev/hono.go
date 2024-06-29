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

func (ctx *Context) ParseBody() ([]byte, error) {
	return io.ReadAll(ctx.ReadBody())
}

func (ctx *Context) ParseText() (string, error) {
	body, err := ctx.ParseBody()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (ctx *Context) ParseJSON(target any) error {
	body, err := ctx.ParseBody()
	if err != nil {
		return err
	}
	return honojson.ParseJSON(string(body), &target)
}

func (ctx *Context) GetParam(key string) string {
	value, ok := ctx.Params[key]
	if !ok {
		return fmt.Sprintf("param '%s' not found", key)
	}
	return value
}

func (ctx *Context) Method() string {
	return ctx.Req.Method
}

func (ctx *Context) Query() url.Values {
	return ctx.Req.URL.Query()
}

func (ctx *Context) Host() string {
	return ctx.Req.Host
}

func (ctx *Context) URL() *url.URL {
	return ctx.Req.URL
}

func (ctx *Context) Proto() string {
	return ctx.Req.Proto
}

func (ctx *Context) ProtoMajor() int {
	return ctx.Req.ProtoMajor
}

func (ctx *Context) ProtoMinor() int {
	return ctx.Req.ProtoMajor
}

func (ctx *Context) ReadBody() io.ReadCloser {
	return ctx.Req.Body
}

func (ctx *Context) GetBody() func() (io.ReadCloser, error) {
	return ctx.Req.GetBody
}

func (ctx *Context) ContentLength() int64 {
	return ctx.Req.ContentLength
}

func (ctx *Context) TransferEncoding() []string {
	return ctx.Req.TransferEncoding
}

func (ctx *Context) Close() bool {
	return ctx.Req.Close
}

func (ctx *Context) Form() url.Values {
	return ctx.Req.Form
}

func (ctx *Context) PostForm() url.Values {
	return ctx.Req.PostForm
}

func (ctx *Context) MultipartForm() *multipart.Form {
	return ctx.Req.MultipartForm
}

func (ctx *Context) RemoteAddr() string {
	return ctx.Req.RemoteAddr
}

func (ctx *Context) ReqURI() string {
	return ctx.Req.RequestURI
}

func (ctx *Context) TLS() *tls.ConnectionState {
	return ctx.Req.TLS
}

func (ctx *Context) Cancel() <-chan struct{} {
	return ctx.Req.Cancel
}

func (ctx *Context) Response() *http.Response {
	return ctx.Req.Response
}
