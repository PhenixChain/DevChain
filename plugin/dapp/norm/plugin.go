// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package norm

import (
	"github.com/PhenixChain/devchain/pluginmgr"
	"github.com/PhenixChain/devchain/plugin/dapp/norm/executor"
	"github.com/PhenixChain/devchain/plugin/dapp/norm/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.NormX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      nil,
		RPC:      nil,
	})
}
