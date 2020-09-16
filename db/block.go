package db

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

const BlockReward = 100

type Block struct {
	Header BlockHeader `json:"header"`
	TXs    []TX        `json:"txs"`
}

type BlockHeader struct {
	Hash   string   `json:"hash"`
	Parent string   `json:"parent_hash"`
	Number uint64   `json:"height"`
	TXs    []string `json:"tx_hashes"`

	Nonce     uint64 `json:"nonce"`
	Timestamp uint64 `json:"timestamp"`
}

func NewBlock(block *Block) error {
	block.Header.Timestamp = uint64(time.Now().Unix())

	hash, err := block.sha256()
	if err != nil {
		return err
	}
	block.Header.Hash = fmt.Sprintf("%x", hash)

	return nil
}

func (b *Block) sha256() ([sha256.Size]byte, error) {

	data, err := json.Marshal(b)
	if err != nil {
		return [sha256.Size]byte{}, err
	}
	return sha256.Sum256(data), nil
}

func (b *Block) IsValid() bool {
	if b.Header.Hash == "" {
		return false
	}
	return true
	// return fmt.Sprintf("%x", hash[0]) == "0" &&
	// 	fmt.Sprintf("%x", hash[1]) == "0" &&
	// 	fmt.Sprintf("%x", hash[2]) == "0" &&
	// 	fmt.Sprintf("%x", hash[3]) != "0"
}

func NewGensisBlock() (*Block, error) {
	coinbase := &TX{
		Timestamp: time.Date(2009, time.January, 3, 0, 0, 0, 0, time.UTC).Unix(),
		Data:      "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks",
	}
	if err := NewTX(coinbase); err != nil {
		return nil, err
	}

	block := &Block{
		Header: BlockHeader{
			Number: 0,
			TXs:    []string{coinbase.Hash},
		},
		TXs: []TX{*coinbase},
	}
	if err := NewBlock(block); err != nil {
		return nil, err
	}

	return block, nil
}
