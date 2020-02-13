package template

import (
	"bytes"
	"errors"
	"io"
	"log"
	"path"
	"strings"
	"text/template"
)

type Templator struct {
	context *template.Template
}

func New(filename string) (Templator, error) {
	tmpl, err := buildTemplate(filename)
	return Templator{context: tmpl}, err
}

func (t Templator) Decode(io.Reader) (interface{}, error) {
	return nil, errors.New("unsupport to decode template")
}

func (t Templator) Encode(out io.Writer, data interface{}) error {
	return t.context.Execute(out, data)
}

func buildTemplate(filename string) (tmpl *template.Template, err error) {
	name := path.Base(filename)

	return template.New(name).Funcs(template.FuncMap{
		"split":   strings.Split,
		"join":    strings.Join,
		"replace": strings.Replace,
		"base":    path.Base,
		"dir":     path.Dir,
		"partial": renderPartial,
	}).ParseFiles(filename)
}

func renderPartial(path string, data interface{}) string {
	tmpl, err := buildTemplate(path)
	if err != nil {
		log.Panicf("failed to build template: %s", err)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)

	if err != nil {
		log.Panicf("failed to build template: %s", err)
	}

	return buf.String()
}
