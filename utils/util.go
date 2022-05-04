package utils

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/render"
)

func Newtemplate(core *kyoto.Core, page, path string) *template.Template {
	return template.Must(template.New(page).Funcs(render.FuncMap(core)).ParseGlob(path))
}
