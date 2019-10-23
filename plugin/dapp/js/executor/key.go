package executor

import (
	"github.com/PhenixChain/devchain/types"
	ptypes "github.com/PhenixChain/devchain/plugin/dapp/js/types"
)

func calcAllPrefix(name string) ([]byte, []byte) {
	execer := types.ExecName("user." + ptypes.JsX + "." + name)
	state := types.CalcStatePrefix([]byte(execer))
	local := types.CalcLocalPrefix([]byte(execer))
	return state, local
}

func calcCodeKey(name string) []byte {
	return append([]byte("mavl-"+ptypes.JsX+"-code-"), []byte(name)...)
}
