package multisig

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	_ "github.com/PhenixChain/devchain/plugin/dapp/multisig/autotest" //register auto test
	"github.com/PhenixChain/devchain/plugin/dapp/multisig/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/multisig/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/multisig/rpc"
	mty "github.com/PhenixChain/devchain/plugin/dapp/multisig/types"
	_ "github.com/PhenixChain/devchain/plugin/dapp/multisig/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     mty.MultiSigX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.MultiSigCmd,
		RPC:      rpc.Init,
	})
}
