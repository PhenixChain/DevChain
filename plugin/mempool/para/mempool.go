package para

import (
	"github.com/PhenixChain/devchain/queue"
	drivers "github.com/PhenixChain/devchain/system/mempool"
	"github.com/PhenixChain/devchain/types"
)

//--------------------------------------------------------------------------------
// Module Mempool

func init() {
	drivers.Reg("para", New)
}

//New 创建price cache 结构的 mempool
func New(cfg *types.Mempool, sub []byte) queue.Module {
	return NewMempool(cfg)
}
