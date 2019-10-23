// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package privacy

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	_ "github.com/PhenixChain/devchain/plugin/dapp/privacy/autotest" // register autotest package
	"github.com/PhenixChain/devchain/plugin/dapp/privacy/commands"
	_ "github.com/PhenixChain/devchain/plugin/dapp/privacy/crypto" // register crypto package
	"github.com/PhenixChain/devchain/plugin/dapp/privacy/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/privacy/rpc"
	"github.com/PhenixChain/devchain/plugin/dapp/privacy/types"
	_ "github.com/PhenixChain/devchain/plugin/dapp/privacy/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.PrivacyX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.PrivacyCmd,
		RPC:      rpc.Init,
	})
}
