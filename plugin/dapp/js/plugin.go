package js

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	"github.com/PhenixChain/devchain/plugin/dapp/js/executor"
	ptypes "github.com/PhenixChain/devchain/plugin/dapp/js/types"

	// init auto test
	_ "github.com/PhenixChain/devchain/plugin/dapp/js/autotest"
	"github.com/PhenixChain/devchain/plugin/dapp/js/command"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     ptypes.JsX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      command.JavaScriptCmd,
		RPC:      nil,
	})
}
