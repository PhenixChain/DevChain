// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

//only load all plugin and system
import (
	"testing"

	rpctypes "github.com/PhenixChain/devchain/rpc/types"
	_ "github.com/PhenixChain/devchain/system"
	_ "github.com/PhenixChain/devchain/plugin"
	ty "github.com/PhenixChain/devchain/plugin/dapp/ticket/types"
	"github.com/stretchr/testify/assert"
)

func TestNewTicket(t *testing.T) {
	//选票(可以用hotwallet 关闭选票)
	in := &ty.TicketClose{MinerAddress: mocker.GetHotAddress()}
	var res rpctypes.ReplyHashes
	err := mocker.GetJSONC().Call("ticket.CloseTickets", in, &res)
	assert.Nil(t, err)
}
