package html_test

import (
	"testing"

	"github.com/evex-dev/hono.go/src/html"
)

func TestRenderHtml(t *testing.T) {
	htmlObject := html.Html(nil, *html.Head(nil, *html.Title(nil, *html.Text("hello"))), *html.Body(nil, *html.H1(nil, *html.Text("hello")), *html.Br(nil)))

	t.Log(html.Render(htmlObject))
}
