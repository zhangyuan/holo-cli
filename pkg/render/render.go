package render

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
)

type Render struct {
	tmpl *template.Template
}

func NewRender(name string, text string) (*Render, error) {
	t := template.New(name).Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap{
		"isLast": IsLast,
	})

	tmpl, err := t.Parse(text)
	if err != nil {
		return nil, err
	}
	return &Render{
		tmpl: tmpl,
	}, nil
}

func (r *Render) Render(data interface{}) (string, error) {
	var buf bytes.Buffer
	if err := r.tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func IsLast(index, length int) bool {
	return index == length-1
}
