package template

import (
	"bytes"
	"github.com/mijime/merje/remarshal"
	"os"
	"path"
	"strings"
	"text/template"
)

type factory struct{}

type converter struct {
	templatePath string
}

func init() {
	remarshal.Factory.Regist("template", factory{})
}

func new(options remarshal.Options) converter {
	return converter{templatePath: options.Format}
}

func (f factory) Lookup(options interface{}) interface{} {
	op, ok := options.(remarshal.Options)

	if !ok {
		return nil
	}

	_, err := os.Stat(op.Format)

	if err != nil {
		return nil
	}

	return new(op)
}

func (c converter) Unmarshal(buf []byte) (data interface{}, err error) {
	return nil, nil
}

func (c converter) Marshal(data interface{}) (output []byte, err error) {
	var (
		buf  bytes.Buffer
		tmpl *template.Template
	)

	tmpl, err = buildTemplate(c.templatePath)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&buf, data)
	return buf.Bytes(), err
}

func buildTemplate(format string) (tmpl *template.Template, err error) {
	funcMap := template.FuncMap{
		"split":   strings.Split,
		"join":    strings.Join,
		"replace": strings.Replace,
		"base":    path.Base,
		"dir":     path.Dir,
	}

	name := path.Base(format)
	return template.New(name).Funcs(funcMap).ParseFiles(format)
}
