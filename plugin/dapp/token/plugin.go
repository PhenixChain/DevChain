// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token 创建token
package token

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	_ "github.com/PhenixChain/devchain/plugin/dapp/token/autotest" // register token autotest package
	"github.com/PhenixChain/devchain/plugin/dapp/token/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/token/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/token/rpc"
	"github.com/PhenixChain/devchain/plugin/dapp/token/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TokenX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TokenCmd,
		RPC:      rpc.Init,
	})
}
