// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echo

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	"github.com/PhenixChain/devchain/plugin/dapp/echo/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/echo/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/echo/rpc"
	echotypes "github.com/PhenixChain/devchain/plugin/dapp/echo/types/echo"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     echotypes.EchoX,
		ExecName: echotypes.EchoX,
		Exec:     executor.Init,
		Cmd:      commands.EchoCmd,
		RPC:      rpc.Init,
	})
}
