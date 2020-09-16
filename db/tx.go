package db

import (
	"crypto/sha256"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type TX struct {
	Hash  Hash           `json:"hash"`
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Value uint           `json:"value"`

	Nonce     uint   `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
}

func NewTX(tx *TX) error {
	tx.Timestamp = time.Now().Unix()

	hash, err := tx.sha256()
	if err != nil {
		return err
	}
	tx.Hash = hash

	return nil
}

func (t TX) sha256() (Hash, error) {
	if (t.Hash != Hash{}) {
		return t.Hash, nil
	}

	txJson, err := json.Marshal(t)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(txJson), nil
}
