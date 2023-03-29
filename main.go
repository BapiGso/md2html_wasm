package main

import (
	"bytes"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"syscall/js"
)

var md = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
		//mathjax.MathJax,
		//&mermaid.Extender{},
		//latex.NewLatex(),
		highlighting.NewHighlighting(
			highlighting.WithStyle("github")),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
		html.WithUnsafe(),
	),
)

func main() {
	js.Global().Set("convertMDToHTML", js.FuncOf(convertMDToHTML))
	select {}
}

func convertMDToHTML(this js.Value, args []js.Value) interface{} {
	source := []byte(args[0].String())
	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		return nil
	}
	return buf.String()
}
