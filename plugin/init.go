package plugin

import (
	_ "github.com/PhenixChain/devchain/plugin/consensus/init" //consensus init
	_ "github.com/PhenixChain/devchain/plugin/crypto/init"    //crypto init
	_ "github.com/PhenixChain/devchain/plugin/dapp/init"      //dapp init
	_ "github.com/PhenixChain/devchain/plugin/mempool/init"   //mempool init
	_ "github.com/PhenixChain/devchain/plugin/store/init"     //store init
)
