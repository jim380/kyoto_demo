package page_index_components

import (
	"encoding/json"
	"net/http"

	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/lifecycle"
)

func ComponentUUID(core *kyoto.Core) {
	lifecycle.Init(core, func() {
		core.State.Set("UUID", "")
	})
	lifecycle.Async(core, func() error {
		resp, _ := http.Get("http://httpbin.org/uuid")
		data := map[string]string{}
		json.NewDecoder(resp.Body).Decode(&data)
		core.State.Set("UUID", data["uuid"])
		return nil
	})
}
