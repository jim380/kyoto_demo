package page_test

import (
	"html/template"

	"github.com/jim380/kyoto-demo/utils"
	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/lifecycle"
	"github.com/kyoto-framework/kyoto/render"

	page_test_components "github.com/jim380/kyoto-demo/pages/test/components"
)

func PageTest(core *kyoto.Core) {
	lifecycle.Init(core, func() {
		// (component_alias, component)
		core.Component("Test", page_test_components.ComponentTest)
	})
	render.Template(core, func() *template.Template {
		t := utils.Newtemplate(core, "page.test.html", "pages/test/*.html")
		return template.Must(t.ParseGlob("pages/test/components/*.html"))
	})
}
