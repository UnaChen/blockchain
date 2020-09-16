package db

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type TX struct {
	Hash  string         `json:"hash"`
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
	tx.Hash = fmt.Sprintf("%x", hash)

	return nil
}

func (t TX) sha256() ([sha256.Size]byte, error) {

	txJson, err := json.Marshal(t)
	if err != nil {
		return [sha256.Size]byte{}, err
	}
	return sha256.Sum256(txJson), nil
}
