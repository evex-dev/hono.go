package html

import "strings"

type HtmlObject struct {
	tagName    string
	attributes map[string]string
	children   []HtmlObject
}

func Render(obj *HtmlObject) string {
	if obj == nil {
		return ""
	}

	if obj.tagName == "$$" {
		return obj.attributes["content"]
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
			body += Render(&v)
		}
		return head + body + "</" + obj.tagName + ">"
	}

	return head + "/>"
}

func Create(tagName string, attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return &HtmlObject{
		tagName:    tagName,
		attributes: attributes,
		children:   children,
	}
}

func A(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("a", attributes, children...)
}

func Abbr(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("abbr", attributes, children...)
}

func Address(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("address", attributes, children...)
}

func Area(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("area", attributes, children...)
}

func Article(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("article", attributes, children...)
}

func Aside(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("aside", attributes, children...)
}

func Audio(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("audio", attributes, children...)
}

func B(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("b", attributes, children...)
}

func Base(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("base", attributes, children...)
}

func Bdi(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("bdi", attributes, children...)
}

func Bdo(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("bdo", attributes, children...)
}

func Big(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("big", attributes, children...)
}

func Blockquote(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("blockquote", attributes, children...)
}

func Body(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("body", attributes, children...)
}

func Br(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("br", attributes, children...)
}

func Button(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("button", attributes, children...)
}

func Canvas(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("canvas", attributes, children...)
}

func Caption(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("caption", attributes, children...)
}

func Center(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("center", attributes, children...)
}

func Cite(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("cite", attributes, children...)
}

func Code(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("code", attributes, children...)
}

func Col(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("col", attributes, children...)
}

func Colgroup(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("colgroup", attributes, children...)
}

func Data(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("data", attributes, children...)
}

func Datalist(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("datalist", attributes, children...)
}

func Dd(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("dd", attributes, children...)
}

func Del(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("del", attributes, children...)
}

func Details(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("details", attributes, children...)
}

func Dfn(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("dfn", attributes, children...)
}

func Dialog(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("dialog", attributes, children...)
}

func Div(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("div", attributes, children...)
}

func Dl(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("dl", attributes, children...)
}

func Dt(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("dt", attributes, children...)
}

func Em(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("em", attributes, children...)
}

func Embed(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("embed", attributes, children...)
}

func Fieldset(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("fieldset", attributes, children...)
}

func Figcaption(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("figcaption", attributes, children...)
}

func Figure(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("figure", attributes, children...)
}

func Footer(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("footer", attributes, children...)
}

func Form(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("form", attributes, children...)
}

func H1(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("h1", attributes, children...)
}

func H2(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("h2", attributes, children...)
}

func H3(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("h3", attributes, children...)
}

func H4(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("h4", attributes, children...)
}

func H5(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("h5", attributes, children...)
}

func H6(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("h6", attributes, children...)
}

func Head(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("head", attributes, children...)
}

func Header(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("header", attributes, children...)
}

func Hr(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("hr", attributes, children...)
}

func Html(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("html", attributes, children...)
}

func I(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("i", attributes, children...)
}

func Iframe(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("iframe", attributes, children...)
}

func Img(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("img", attributes, children...)
}

func Input(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("input", attributes, children...)
}

func Ins(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("ins", attributes, children...)
}

func Kbd(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("kbd", attributes, children...)
}

func Label(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("label", attributes, children...)
}

func Legend(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("legend", attributes, children...)
}

func Li(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("li", attributes, children...)
}

func Link(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("link", attributes, children...)
}

func Main(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("main", attributes, children...)
}

func Map(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("map", attributes, children...)
}

func Menu(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("menu", attributes, children...)
}

func Meta(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("meta", attributes, children...)
}

func Nav(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("nav", attributes, children...)
}

func Noscript(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("noscript", attributes, children...)
}

func Ol(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("ol", attributes, children...)
}

func Optgroup(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("optgroup", attributes, children...)
}

func Option(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("option", attributes, children...)
}

func Output(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("output", attributes, children...)
}

func P(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("p", attributes, children...)
}

func Param(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("param", attributes, children...)
}

func Pre(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("pre", attributes, children...)
}

func Progress(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("progress", attributes, children...)
}

func Q(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("q", attributes, children...)
}

func Rp(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("rp", attributes, children...)
}

func Rt(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("rt", attributes, children...)
}

func Ruby(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("ruby", attributes, children...)
}

func S(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("s", attributes, children...)
}

func Samp(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("samp", attributes, children...)
}

func Script(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("script", attributes, children...)
}

func Section(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("section", attributes, children...)
}

func Select(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("select", attributes, children...)
}

func Small(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("small", attributes, children...)
}

func Source(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("source", attributes, children...)
}

func Span(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("span", attributes, children...)
}

func Strong(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("strong", attributes, children...)
}

func Style(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("style", attributes, children...)
}

func Sub(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("sub", attributes, children...)
}

func Summary(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("summary", attributes, children...)
}

func Sup(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("sup", attributes, children...)
}

func Table(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("table", attributes, children...)
}

func Tbody(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("tbody", attributes, children...)
}

func Td(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("td", attributes, children...)
}

func Template(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("template", attributes, children...)
}

func Textarea(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("textarea", attributes, children...)
}

func Tfoot(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("tfoot", attributes, children...)
}

func Th(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("th", attributes, children...)
}

func Thead(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("thead", attributes, children...)
}

func Time(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("time", attributes, children...)
}

func Title(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("title", attributes, children...)
}

func Tr(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("tr", attributes, children...)
}

func Track(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("track", attributes, children...)
}

func U(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("u", attributes, children...)
}

func Ul(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("ul", attributes, children...)
}

func Var(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("var", attributes, children...)
}

func Video(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("video", attributes, children...)
}

func Wbr(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("wbr", attributes, children...)
}

func Xmp(attributes map[string]string, children ...HtmlObject) *HtmlObject {
	return Create("xmp", attributes, children...)
}

func Text(children ...string) *HtmlObject {
	return Create("$$", map[string]string{
		"content": EscapeHTML(strings.Join(children, "")),
	})
}

func Raw(children ...string) *HtmlObject {
	return Create("$$", map[string]string{
		"content": strings.Join(children, ""),
	})
}