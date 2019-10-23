// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"github.com/PhenixChain/devchain/common/log/log15"
	"github.com/PhenixChain/devchain/rpc/types"
	pty "github.com/PhenixChain/devchain/plugin/dapp/privacy/types"
)

var log = log15.New("module", "privacy.rpc")

// Jrpc json rpc class
type Jrpc struct {
	cli *channelClient
}

// Grpc grpc class
type Grpc struct {
	*channelClient
}

type channelClient struct {
	types.ChannelClient
}

// Init init rpc server
func Init(name string, s types.RPCServer) {
	cli := &channelClient{}
	grpc := &Grpc{channelClient: cli}
	cli.Init(name, s, &Jrpc{cli: cli}, grpc)
	pty.RegisterPrivacyServer(s.GRPC(), grpc)
}
