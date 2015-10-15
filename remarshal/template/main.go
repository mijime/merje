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
	remarshal.Factory.Regist("template", &factory{})
}

func New(option remarshal.Option) *converter {
	return &converter{option.Format}
}

func (this *factory) Lookup(option interface{}) interface{} {
	roption, ok := option.(remarshal.Option)

	if !ok {
		return nil
	}

	_, err := os.Stat(roption.Format)

	if err != nil {
		return nil
	}

	return New(roption)
}

func (this *converter) Unmarshal(buf []byte) (data interface{}, err error) {
	return nil, nil
}

func (this *converter) Marshal(data interface{}) (output []byte, err error) {
	var (
		buf  bytes.Buffer
		tmpl *template.Template
	)

	tmpl, err = buildTemplate(this.templatePath)
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
