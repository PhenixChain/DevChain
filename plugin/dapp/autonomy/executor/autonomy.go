// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"github.com/PhenixChain/devchain/common/address"
	log "github.com/PhenixChain/devchain/common/log/log15"
	drivers "github.com/PhenixChain/devchain/system/dapp"
	"github.com/PhenixChain/devchain/types"
	auty "github.com/PhenixChain/devchain/plugin/dapp/autonomy/types"
)

type subConfig struct {
	Total      string `json:"total"`
	UseBalance bool   `json:"useBalance"`
}

var (
	alog         = log.New("module", "execs.autonomy")
	driverName   = auty.AutonomyX
	autonomyAddr = address.ExecAddress(auty.AutonomyX)
	cfg          subConfig
)

func init() {
	ety := types.LoadExecutorType(driverName)
	ety.InitFuncList(types.ListMethod(&Autonomy{}))
}

// Init 重命名执行器名称
func Init(name string, sub []byte) {
	if sub != nil {
		types.MustDecode(sub, &cfg)
	}
	drivers.Register(GetName(), newAutonomy, types.GetDappFork(driverName, "Enable"))
}

// Autonomy 执行器结构体
type Autonomy struct {
	drivers.DriverBase
}

func newAutonomy() drivers.Driver {
	t := &Autonomy{}
	t.SetChild(t)
	t.SetExecutorType(types.LoadExecutorType(driverName))
	return t
}

// GetName 获得执行器名字
func GetName() string {
	return newAutonomy().GetName()
}

// GetDriverName 获得驱动名字
func (u *Autonomy) GetDriverName() string {
	return driverName
}
