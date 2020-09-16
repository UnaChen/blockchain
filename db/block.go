package db

import (
	"crypto/sha256"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const BlockReward = 100

// func (h Hash) MarshalText() ([]byte, error) {
// 	return []byte(h.Hex()), nil
// }

// func (h *Hash) UnmarshalText(data []byte) error {
// 	_, err := hex.Decode(h[:], data)
// 	return err
// }

// func (h Hash) Hex() string {
// 	return hex.EncodeToString(h[:])
// }

// func (h Hash) IsEmpty() bool {
// 	emptyHash := Hash{}

// 	return bytes.Equal(emptyHash[:], h[:])
// }

type Block struct {
	Header BlockHeader `json:"header"`
	TXs    []TX        `json:"transcations"`
}

type BlockHeader struct {
	Hash   Hash   `json:"block_hash"`
	Parent Hash   `json:"parent_hash"`
	Number uint64 `json:"height"`
	TXs    []Hash `json:"txs_hash"`

	Miner     common.Address `json:"miner"`
	Nonce     uint64         `json:"nonce"`
	Timestamp int64          `json:"timestamp"`
}

func NewBlock(block *Block) error {
	block.Header.Timestamp = time.Now().Unix()

	hash, err := block.sha256()
	if err != nil {
		return err
	}
	block.Header.Hash = hash

	return nil
}

func (b *Block) sha256() (Hash, error) {
	if (b.Header.Hash != Hash{}) {
		return b.Header.Hash, nil
	}

	data, err := json.Marshal(b)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(data), nil
}

func (b *Block) IsValid() bool {
	if (b.Header.Hash == Hash{}) {
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
			TXs:    []Hash{coinbase.Hash},
		},
		TXs: []TX{*coinbase},
	}
	if err := NewBlock(block); err != nil {
		return nil, err
	}

	return block, nil
}
