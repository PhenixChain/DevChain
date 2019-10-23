// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package system 系统基础插件包
package system

import (
	_ "github.com/PhenixChain/devchain/system/consensus/init" //register consensus init package
	_ "github.com/PhenixChain/devchain/system/crypto/init"
	_ "github.com/PhenixChain/devchain/system/dapp/init"
	_ "github.com/PhenixChain/devchain/system/mempool/init"
	_ "github.com/PhenixChain/devchain/system/store/init"
)
