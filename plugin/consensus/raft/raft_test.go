// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raft

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/PhenixChain/devchain/blockchain"
	"github.com/PhenixChain/devchain/common"
	"github.com/PhenixChain/devchain/common/address"
	"github.com/PhenixChain/devchain/common/crypto"
	"github.com/PhenixChain/devchain/common/limits"
	"github.com/PhenixChain/devchain/common/log"
	"github.com/PhenixChain/devchain/executor"
	"github.com/PhenixChain/devchain/mempool"
	"github.com/PhenixChain/devchain/p2p"
	"github.com/PhenixChain/devchain/queue"
	"github.com/PhenixChain/devchain/store"
	"github.com/PhenixChain/devchain/types"

	_ "github.com/PhenixChain/devchain/system"
	_ "github.com/PhenixChain/devchain/plugin/dapp/init"
	pty "github.com/PhenixChain/devchain/plugin/dapp/norm/types"
	_ "github.com/PhenixChain/devchain/plugin/store/init"
)

var (
	random    *rand.Rand
	txNumber  = 10
	loopCount = 10
)

func init() {
	err := limits.SetLimits()
	if err != nil {
		panic(err)
	}
	random = rand.New(rand.NewSource(types.Now().UnixNano()))
	log.SetLogLevel("info")
}
func TestRaftPerf(t *testing.T) {
	RaftPerf()
	fmt.Println("=======start clear test data!=======")
	clearTestData()
}
func RaftPerf() {
	q, chain, s, mem, exec, cs, p2p := initEnvRaft()
	defer q.Close()
	defer s.Close()
	defer p2p.Close()
	defer mem.Close()
	defer exec.Close()
	defer chain.Close()
	defer cs.Close()

	sendReplyList(q)
}

func initEnvRaft() (queue.Queue, *blockchain.BlockChain, queue.Module, queue.Module, *executor.Executor, queue.Module, queue.Module) {
	var q = queue.New("channel")
	flag.Parse()
	cfg, sub := types.InitCfg("chain33.test.toml")
	types.Init(cfg.Title, cfg)

	s := store.New(cfg.Store, sub.Store)
	s.SetQueueClient(q.Client())

	chain := blockchain.New(cfg.BlockChain)
	chain.SetQueueClient(q.Client())

	exec := executor.New(cfg.Exec, sub.Exec)
	exec.SetQueueClient(q.Client())
	types.SetMinFee(0)

	mem := mempool.New(cfg.Mempool, nil)
	mem.SetQueueClient(q.Client())

	cs := NewRaftCluster(cfg.Consensus, sub.Consensus["raft"])
	cs.SetQueueClient(q.Client())

	network := p2p.New(cfg.P2P)

	network.SetQueueClient(q.Client())
	return q, chain, s, mem, exec, cs, network
}

func generateKey(i, valI int) string {
	key := make([]byte, valI)
	binary.PutUvarint(key[:10], uint64(valI))
	binary.PutUvarint(key[12:24], uint64(i))
	if _, err := rand.Read(key[24:]); err != nil {
		os.Exit(1)
	}
	return string(key)
}

func generateValue(i, valI int) string {
	value := make([]byte, valI)
	binary.PutUvarint(value[:16], uint64(i))
	binary.PutUvarint(value[32:128], uint64(i))
	if _, err := rand.Read(value[128:]); err != nil {
		os.Exit(1)
	}
	return string(value)
}

func getprivkey(key string) crypto.PrivKey {
	cr, err := crypto.New(types.GetSignName("", types.SECP256K1))
	if err != nil {
		panic(err)
	}
	bkey, err := common.FromHex(key)
	if err != nil {
		panic(err)
	}
	priv, err := cr.PrivKeyFromBytes(bkey)
	if err != nil {
		panic(err)
	}
	return priv
}

func sendReplyList(q queue.Queue) {
	client := q.Client()
	client.Sub("mempool")
	var count int
	for msg := range client.Recv() {
		if msg.Ty == types.EventTxList {
			count++
			msg.Reply(client.NewMessage("consensus", types.EventReplyTxList,
				&types.ReplyTxList{Txs: getReplyList(txNumber)}))
			if count >= loopCount {
				time.Sleep(4 * time.Second)
				break
			}
		}
	}
}

func prepareTxList() *types.Transaction {
	var key string
	var value string
	var i int

	key = generateKey(i, 32)
	value = generateValue(i, 180)

	nput := &pty.NormAction_Nput{Nput: &pty.NormPut{Key: []byte(key), Value: []byte(value)}}
	action := &pty.NormAction{Value: nput, Ty: pty.NormActionPut}
	tx := &types.Transaction{Execer: []byte("norm"), Payload: types.Encode(action), Fee: 0}
	tx.To = address.ExecAddress("norm")
	tx.Nonce = random.Int63()
	tx.Sign(types.SECP256K1, getprivkey("CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"))
	return tx
}

func getReplyList(n int) (txs []*types.Transaction) {

	for i := 0; i < n; i++ {
		txs = append(txs, prepareTxList())
	}
	return txs
}

func clearTestData() {
	err := os.RemoveAll("datadir")
	if err != nil {
		fmt.Println("delete datadir have a err:", err.Error())
	}
	err = os.RemoveAll("chain33_raft-1")
	if err != nil {
		fmt.Println("delete chain33_raft dir have a err:", err.Error())
	}
	fmt.Println("test data clear successfully!")
}
