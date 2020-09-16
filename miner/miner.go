package miner

import (
	"blockchain/db"
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const IntervalSeconds = 10

type Miner struct{}

func (m *Miner) Mine(ctx context.Context, block *db.Block) error {
	if len(block.TXs) == 0 {
		return fmt.Errorf("mining empty blocks is not allowed")
	}

	var (
		start   = time.Now()
		attempt = 0
	)
	block.Header.Nonce = generateNonce()

	for !block.IsValid() {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "mining cancelled")
		default:

			attempt++
			block.Header.Nonce = generateNonce()

			if attempt%1000000 == 0 || attempt == 1 {
				logrus.WithField("# txs", len(block.TXs)).WithField("# attempts", attempt).Infoln("mining ...")
			}

			if err := db.NewBlock(block); err != nil {
				return errors.Wrap(err, "failt to hash block")
			}
		}
	}

	fmt.Printf("\nMined new Block '%s' using PoW >v<:\n", block.Header.Hash)
	fmt.Printf("\tHeight: '%v'\n", block.Header.Number)
	fmt.Printf("\tParent: '%x'\n\n", block.Header.Parent)
	fmt.Printf("\tNonce: '%v'\n", block.Header.Nonce)
	fmt.Printf("\tCreated: '%v'\n", block.Header.Timestamp)
	fmt.Printf("\tTXs: '%x'\n", block.Header.TXs)

	fmt.Printf("\tAttempt: '%v'\n", attempt)
	fmt.Printf("\tSpent Time: %s\n\n", time.Since(start))

	return nil
}

func generateNonce() uint64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Uint64()
}
