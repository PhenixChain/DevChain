// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paracross

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	_ "github.com/PhenixChain/devchain/plugin/dapp/paracross/autotest" // register autotest package
	"github.com/PhenixChain/devchain/plugin/dapp/paracross/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/paracross/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/paracross/rpc"
	"github.com/PhenixChain/devchain/plugin/dapp/paracross/types"
	_ "github.com/PhenixChain/devchain/plugin/dapp/paracross/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ParaX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.ParcCmd,
		RPC:      rpc.Init,
	})
}
