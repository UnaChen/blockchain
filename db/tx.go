package db

import (
	"crypto/sha256"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// func NewAccount(value string) common.Address {
// 	return common.HexToAddress(value)
// }

type Tx struct {
	Hash      Hash           `json:"hash"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Value     uint           `json:"value"`
	Timestamp int64          `json:"timestamp"`
}

func NewTx(from, to common.Address, value uint) (*Tx, error) {
	tx := &Tx{
		From:      from,
		To:        to,
		Value:     value,
		Timestamp: time.Now().Unix(),
	}

	hash, err := tx.sha256()
	if err != nil {
		return nil, err
	}
	tx.Hash = hash

	return tx, nil
}

func (t Tx) sha256() (Hash, error) {
	txJson, err := json.Marshal(t)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(txJson), nil
}
