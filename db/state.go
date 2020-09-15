package db

import (
	"github.com/ethereum/go-ethereum/common"
)

type State struct {
	Balances map[common.Address]uint

	latestBlock     Block
	latestBlockHash Hash
	hasGenesisBlock bool
}
