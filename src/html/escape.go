package html

import "strings"

var htmlReplacer = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	"\"", "&quot;",
	"'", "&#39;",
	"\"", "&quot;",
)

func EscapeHTML(s string) string {
	return htmlReplacer.Replace(s)
}
