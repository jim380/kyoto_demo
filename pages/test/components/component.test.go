package page_test_components

import (
	"encoding/json"
	"strings"

	"github.com/jim380/kyoto-demo/queries"
	"github.com/kyoto-framework/kyoto"
	"github.com/kyoto-framework/kyoto/lifecycle"
	"go.uber.org/zap"
)

type nodeStatus struct {
	result `json:"result"`
}

type result struct {
	nodeInfo `json:"node_info"`
}

type nodeInfo struct {
	Network string `json:"network"`
	Moniker string `json:"moniker"`
}

func ComponentTest(core *kyoto.Core) {
	lifecycle.Init(core, func() {
		core.State.Set("Network", "")
		core.State.Set("Moniker", "")
	})

	var nodeStatus nodeStatus

	res, err := queries.RESTQuery("/status")
	if err != nil {
		zap.L().Fatal("", zap.Bool("Success", false), zap.String("err", err.Error()))
	}
	json.Unmarshal(res, &nodeStatus)
	if strings.Contains(string(res), "not found") {
		zap.L().Fatal("", zap.Bool("Success", false), zap.String("err", string(res)))
	} else if strings.Contains(string(res), "error:") || strings.Contains(string(res), "error\\\":") {
		zap.L().Fatal("", zap.Bool("Success", false), zap.String("err", string(res)))
	}

	// rd.NodeInfo = nodeInfo
	lifecycle.Async(core, func() error {
		core.State.Set("Network", nodeStatus.Network)
		core.State.Set("Moniker", nodeStatus.Moniker)
		return nil
	})
}
