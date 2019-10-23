// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dposvote

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	"github.com/PhenixChain/devchain/plugin/dapp/dposvote/commands"
	"github.com/PhenixChain/devchain/plugin/dapp/dposvote/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/dposvote/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.DPosX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.DPosCmd,
	})
}
