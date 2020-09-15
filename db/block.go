package db

import (
	"crypto/sha256"
	"encoding/json"
)

// const BlockReward = 100

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
	Txs    []Tx        `json:"payload"`
}

type BlockHeader struct {
	Hash   Hash   `json:"block_hash"`
	Parent Hash   `json:"parent_hash"`
	Number uint64 `json:"height"`
	Txs    []Hash `json:"txs_hash"`

	Nonce uint32 `json:"nonce"`
	Time  uint64 `json:"time"`
}

func NewBlock(header BlockHeader) *Block {
	return &Block{
		Header: header,
	}
}

func (b Block) SHA256() (Hash, error) {
	blockJson, err := json.Marshal(b.Header)
	if err != nil {
		return Hash{}, err
	}

	return sha256.Sum256(blockJson), nil
}

func IsBlockHashValid(hash Hash) bool {
	return true
	// return fmt.Sprintf("%x", hash[0]) == "0" &&
	// 	fmt.Sprintf("%x", hash[1]) == "0" &&
	// 	fmt.Sprintf("%x", hash[2]) == "0" &&
	// 	fmt.Sprintf("%x", hash[3]) != "0"
}
