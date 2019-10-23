// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

import (
	"strings"
	"testing"

	commonlog "github.com/PhenixChain/devchain/common/log"
	"github.com/PhenixChain/devchain/rpc/jsonclient"
	"github.com/PhenixChain/devchain/types"
	"github.com/PhenixChain/devchain/util/testnode"
	"github.com/stretchr/testify/assert"

	_ "github.com/PhenixChain/devchain/system"
	_ "github.com/PhenixChain/devchain/plugin"
)

func init() {
	commonlog.SetLogLevel("error")
}

func TestJRPCChannel(t *testing.T) {
	// 启动RPCmocker
	mocker := testnode.New("--notset--", nil)
	defer func() {
		mocker.Close()
	}()
	mocker.Listen()

	jrpcClient := mocker.GetJSONC()
	assert.NotNil(t, jrpcClient)

	testCases := []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testPropBoardTxCmd},
		{fn: testRevokeProposalBoardTxCmd},
		{fn: testVoteProposalBoardTxCmd},
		{fn: testTerminateProposalBoardTxCmd},
		{fn: testGetProposalBoardCmd},
		{fn: testListProposalBoardCmd},
		{fn: testGetActiveBoardCmd},

		{fn: testPropProjectTxCmd},
		{fn: testRevokeProposalProjectTxCmd},
		{fn: testVoteProposalProjectTxCmd},
		{fn: testPubVoteProposalProjectTxCmd},
		{fn: testTerminateProposalProjectTxCmd},
		{fn: testGetProposalProjectCmd},
		{fn: testListProposalProjectCmd},

		{fn: testPropRuleTxCmd},
		{fn: testRevokeProposalRuleTxCmd},
		{fn: testVoteProposalRuleTxCmd},
		{fn: testTerminateProposalRuleTxCmd},
		{fn: testGetProposalRuleCmd},
		{fn: testListProposalRuleCmd},
		{fn: testGetActiveRuleCmd},

		{fn: testTransferFundTxCmd},
		{fn: testCommentProposalTxCmd},
		{fn: testListProposalCommentCmd},

		{fn: testPropChangeTxCmd},
		{fn: testRevokeProposalChangeTxCmd},
		{fn: testVoteProposalChangeTxCmd},
		{fn: testTerminateProposalChangeTxCmd},
		{fn: testGetProposalChangeCmd},
		{fn: testListProposalChangeCmd},
	}
	for index, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		if err == nil {
			continue
		}
		assert.NotEqualf(t, err, types.ErrActionNotSupport, "test index %d", index)
		if strings.Contains(err.Error(), "rpc: can't find") {
			assert.FailNowf(t, err.Error(), "test index %d", index)
		}
	}
}
