// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/PhenixChain/devchain/cmd/cli/buildflags"
	_ "github.com/PhenixChain/devchain/plugin"
	_ "github.com/PhenixChain/devchain/system"
	"github.com/PhenixChain/devchain/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:8801"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName)
}
