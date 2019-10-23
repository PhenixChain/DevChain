// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package evm

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	"github.com/PhenixChain/devchain/plugin/dapp/evm/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/evm/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/evm/rpc"
	"github.com/PhenixChain/devchain/plugin/dapp/evm/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ExecutorName,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.EvmCmd,
		RPC:      rpc.Init,
	})
}
