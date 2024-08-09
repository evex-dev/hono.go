package html

import "strings"

type Html struct{}

type HtmlObject struct {
	tagName    string
	attributes map[string]string
	children   []string
}

func (h *Html) Render(obj *HtmlObject) string {
	if obj == nil {
		return ""
	}

	head := "<" + obj.tagName
	for k, v := range obj.attributes {
		if v == "" {
			head += " " + k
			continue
		}

		head += " " + EscapeHTML(k) + "=\"" + EscapeHTML(v) + "\""
	}

	if len(obj.children) > 0 {
		head += ">"

		body := ""
		for _, v := range obj.children {
			body += v
		}
		return head + body + "</" + obj.tagName + ">"
	}

	return head + "/>"
}

func (h *Html) Create(tagName string, attributes map[string]string, children ...string) *HtmlObject {
	return &HtmlObject{
		tagName:    tagName,
		attributes: attributes,
		children:   children,
	}
}

func (h *Html) A(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("a", attributes, children...)
}

func (h *Html) Abbr(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("abbr", attributes, children...)
}

func (h *Html) Address(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("address", attributes, children...)
}

func (h *Html) Area(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("area", attributes, children...)
}

func (h *Html) Article(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("article", attributes, children...)
}

func (h *Html) Aside(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("aside", attributes, children...)
}

func (h *Html) Audio(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("audio", attributes, children...)
}

func (h *Html) B(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("b", attributes, children...)
}

func (h *Html) Base(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("base", attributes, children...)
}

func (h *Html) Bdi(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("bdi", attributes, children...)
}

func (h *Html) Bdo(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("bdo", attributes, children...)
}

func (h *Html) Big(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("big", attributes, children...)
}

func (h *Html) Blockquote(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("blockquote", attributes, children...)
}

func (h *Html) Body(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("body", attributes, children...)
}

func (h *Html) Br(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("br", attributes, children...)
}

func (h *Html) Button(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("button", attributes, children...)
}

func (h *Html) Canvas(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("canvas", attributes, children...)
}

func (h *Html) Caption(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("caption", attributes, children...)
}

func (h *Html) Center(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("center", attributes, children...)
}

func (h *Html) Cite(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("cite", attributes, children...)
}

func (h *Html) Code(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("code", attributes, children...)
}

func (h *Html) Col(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("col", attributes, children...)
}

func (h *Html) Colgroup(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("colgroup", attributes, children...)
}

func (h *Html) Data(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("data", attributes, children...)
}

func (h *Html) Datalist(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("datalist", attributes, children...)
}

func (h *Html) Dd(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("dd", attributes, children...)
}

func (h *Html) Del(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("del", attributes, children...)
}

func (h *Html) Details(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("details", attributes, children...)
}

func (h *Html) Dfn(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("dfn", attributes, children...)
}

func (h *Html) Dialog(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("dialog", attributes, children...)
}

func (h *Html) Div(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("div", attributes, children...)
}

func (h *Html) Dl(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("dl", attributes, children...)
}

func (h *Html) Dt(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("dt", attributes, children...)
}

func (h *Html) Em(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("em", attributes, children...)
}

func (h *Html) Embed(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("embed", attributes, children...)
}

func (h *Html) Fieldset(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("fieldset", attributes, children...)
}

func (h *Html) Figcaption(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("figcaption", attributes, children...)
}

func (h *Html) Figure(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("figure", attributes, children...)
}

func (h *Html) Footer(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("footer", attributes, children...)
}

func (h *Html) Form(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("form", attributes, children...)
}

func (h *Html) H1(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("h1", attributes, children...)
}

func (h *Html) H2(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("h2", attributes, children...)
}

func (h *Html) H3(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("h3", attributes, children...)
}

func (h *Html) H4(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("h4", attributes, children...)
}

func (h *Html) H5(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("h5", attributes, children...)
}

func (h *Html) H6(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("h6", attributes, children...)
}

func (h *Html) Head(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("head", attributes, children...)
}

func (h *Html) Header(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("header", attributes, children...)
}

func (h *Html) Hr(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("hr", attributes, children...)
}

func (h *Html) Html(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("html", attributes, children...)
}

func (h *Html) I(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("i", attributes, children...)
}

func (h *Html) Iframe(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("iframe", attributes, children...)
}

func (h *Html) Img(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("img", attributes, children...)
}

func (h *Html) Input(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("input", attributes, children...)
}

func (h *Html) Ins(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("ins", attributes, children...)
}

func (h *Html) Kbd(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("kbd", attributes, children...)
}

func (h *Html) Label(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("label", attributes, children...)
}

func (h *Html) Legend(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("legend", attributes, children...)
}

func (h *Html) Li(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("li", attributes, children...)
}

func (h *Html) Link(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("link", attributes, children...)
}

func (h *Html) Main(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("main", attributes, children...)
}

func (h *Html) Map(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("map", attributes, children...)
}

func (h *Html) Menu(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("menu", attributes, children...)
}

func (h *Html) Meta(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("meta", attributes, children...)
}

func (h *Html) Nav(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("nav", attributes, children...)
}

func (h *Html) Noscript(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("noscript", attributes, children...)
}

func (h *Html) Ol(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("ol", attributes, children...)
}

func (h *Html) Optgroup(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("optgroup", attributes, children...)
}

func (h *Html) Option(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("option", attributes, children...)
}

func (h *Html) Output(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("output", attributes, children...)
}

func (h *Html) P(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("p", attributes, children...)
}

func (h *Html) Param(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("param", attributes, children...)
}

func (h *Html) Pre(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("pre", attributes, children...)
}

func (h *Html) Progress(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("progress", attributes, children...)
}

func (h *Html) Q(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("q", attributes, children...)
}

func (h *Html) Rp(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("rp", attributes, children...)
}

func (h *Html) Rt(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("rt", attributes, children...)
}

func (h *Html) Ruby(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("ruby", attributes, children...)
}

func (h *Html) S(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("s", attributes, children...)
}

func (h *Html) Samp(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("samp", attributes, children...)
}

func (h *Html) Script(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("script", attributes, children...)
}

func (h *Html) Section(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("section", attributes, children...)
}

func (h *Html) Select(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("select", attributes, children...)
}

func (h *Html) Small(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("small", attributes, children...)
}

func (h *Html) Source(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("source", attributes, children...)
}

func (h *Html) Span(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("span", attributes, children...)
}

func (h *Html) Strong(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("strong", attributes, children...)
}

func (h *Html) Style(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("style", attributes, children...)
}

func (h *Html) Sub(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("sub", attributes, children...)
}

func (h *Html) Summary(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("summary", attributes, children...)
}

func (h *Html) Sup(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("sup", attributes, children...)
}

func (h *Html) Table(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("table", attributes, children...)
}

func (h *Html) Tbody(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("tbody", attributes, children...)
}

func (h *Html) Td(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("td", attributes, children...)
}

func (h *Html) Template(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("template", attributes, children...)
}

func (h *Html) Textarea(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("textarea", attributes, children...)
}

func (h *Html) Tfoot(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("tfoot", attributes, children...)
}

func (h *Html) Th(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("th", attributes, children...)
}

func (h *Html) Thead(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("thead", attributes, children...)
}

func (h *Html) Time(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("time", attributes, children...)
}

func (h *Html) Title(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("title", attributes, children...)
}

func (h *Html) Tr(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("tr", attributes, children...)
}

func (h *Html) Track(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("track", attributes, children...)
}

func (h *Html) U(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("u", attributes, children...)
}

func (h *Html) Ul(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("ul", attributes, children...)
}

func (h *Html) Var(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("var", attributes, children...)
}

func (h *Html) Video(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("video", attributes, children...)
}

func (h *Html) Wbr(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("wbr", attributes, children...)
}

func (h *Html) Xmp(attributes map[string]string, children ...string) *HtmlObject {
	return h.Create("xmp", attributes, children...)
}

func (h *Html) Text(text ...string) *HtmlObject {
	return h.Create("$$", map[string]string{
		"content": EscapeHTML(strings.Join(text, "")),
	})
}

func (h *Html) Raw(children ...string) *HtmlObject {
	return h.Create("$$", map[string]string{
		"content": strings.Join(children, ""),
	})
}