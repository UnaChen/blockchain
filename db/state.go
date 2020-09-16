package db

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/ethereum/go-ethereum/common"
)

const DefaultAccountCoin = 100

type State struct {
	LatestBlockHeader BlockHeader
	Account2Nonce     map[common.Address]uint

	blocks   []Block
	balances map[common.Address]uint
}

func NewState(gensis Block) *State {
	return &State{
		LatestBlockHeader: gensis.Header,

		blocks:        []Block{gensis},
		balances:      make(map[common.Address]uint),
		Account2Nonce: make(map[common.Address]uint),
	}

}

func (s *State) AddBlock(b Block) error {
	if _, err := s.isValidBlock(b); err != nil {
		return err
	}

	pendingState := *s
	err := pendingState.addTXs(b.TXs)
	if err != nil {
		return err
	}

	s.LatestBlockHeader = pendingState.LatestBlockHeader

	s.balances = pendingState.balances
	s.Account2Nonce = pendingState.Account2Nonce
	s.blocks = append(s.blocks, b)

	return nil
}

func (s *State) isValidBlock(b Block) (bool, error) {
	if !b.IsValid() {
		return false, fmt.Errorf("invalid block hash %x", b.Header.Hash)
	}

	nextExpectedBlockNumber := s.LatestBlockHeader.Number + 1
	if b.Header.Number != nextExpectedBlockNumber {
		return false, fmt.Errorf("next expected block number must be '%d' not '%d'", nextExpectedBlockNumber, b.Header.Number)
	}

	if s.LatestBlockHeader.Number > 0 && !reflect.DeepEqual(b.Header.Parent, s.LatestBlockHeader.Hash) {
		return false, fmt.Errorf("next block parent hash must be '%x' not '%x'", s.LatestBlockHeader.Hash, b.Header.Parent)
	}

	return true, nil
}

func (s *State) addTXs(txs []TX) error {
	sort.Slice(txs, func(i, j int) bool {
		return txs[i].Timestamp < txs[j].Timestamp
	})
	for _, tx := range txs {
		err := s.addTX(tx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *State) addTX(tx TX) error {

	//credit the new sender or/and the new recipient 100 coins.
	if _, ok := s.balances[tx.From]; !ok {
		s.balances[tx.From] = DefaultAccountCoin
	}
	if _, ok := s.balances[tx.To]; !ok {
		s.balances[tx.To] = DefaultAccountCoin
	}

	expectedNonce := s.Account2Nonce[tx.From] + 1
	if tx.Nonce != expectedNonce {
		return fmt.Errorf("wrong TX. Sender '%s' next nonce must be '%d', not '%d'", tx.From.String(), expectedNonce, tx.Nonce)
	}

	if tx.Value > s.balances[tx.From] {
		return fmt.Errorf("wrong TX. Sender '%s' balance is %d coins. TX cost is %d coins", tx.From.String(), s.balances[tx.From], tx.Value)
	}

	s.balances[tx.From] -= tx.Value
	s.balances[tx.To] += tx.Value

	s.Account2Nonce[tx.From] = tx.Nonce
	return nil
}

func (s *State) GetBalances() map[common.Address]uint {
	return s.balances
}

func (s *State) GetBlocks() []Block {
	return s.blocks
}
