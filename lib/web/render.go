package web

import (
	"html/template"

	"github.com/unrolled/render"
)

var r *render.Render

func init() {
	// setup template rendering
	r = render.New(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs: []template.FuncMap{template.FuncMap{
			"isEvenNumber": isEvenNumber,
			"isOddNumber":  isOddNumber,
		}},
		IndentJSON: true,
	})
}

func isEvenNumber(input int) bool {
	return input%2 == 0
}

func isOddNumber(input int) bool {
	return !isEvenNumber(input)
}

func Render() *render.Render {
	return r
}
