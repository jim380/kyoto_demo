package page_index

import (
	"html/template"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/lifecycle"
	"github.com/kyoto-framework/kyoto/render"

	page_index_components "github.com/jim380/kyoto-demo/pages/index/components"
	"github.com/jim380/kyoto-demo/utils"
)

func PageIndex(core *kyoto.Core) {
	lifecycle.Init(core, func() {
		core.Component("UUID1", page_index_components.ComponentUUID)
		core.Component("UUID2", page_index_components.ComponentUUID)
	})
	render.Template(core, func() *template.Template {
		t := utils.Newtemplate(core, "page.index.html", "pages/index/*.html")
		return template.Must(t.ParseGlob("pages/index/components/*.html"))
	})
}
