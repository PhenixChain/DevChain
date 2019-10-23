// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package guess

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	"github.com/PhenixChain/devchain/plugin/dapp/guess/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/guess/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/guess/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.GuessX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.GuessCmd,
	})
}
