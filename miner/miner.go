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

type Miner struct{}

func (m *Miner) Mine(ctx context.Context, block *db.Block) error {
	if len(block.Txs) == 0 {
		return fmt.Errorf("mining empty blocks is not allowed")
	}

	var (
		start   = time.Now()
		attempt = 0
	)

	var nonce uint32

	for !db.IsBlockHashValid(block.Header.Hash) {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "mining cancelled")
		default:

			attempt++
			nonce = generateNonce()

			if attempt%1000000 == 0 || attempt == 1 {
				logrus.WithField("# txs", len(block.Txs)).WithField("# attempts", attempt).Infoln("mining ...")
			}

			block.Header.Nonce = nonce

			hash, err := block.SHA256()
			if err != nil {
				return errors.Wrap(err, "failt to hash block")
			}

			block.Header.Hash = hash
		}
	}

	fmt.Printf("\nMined new Block '%x' using PoW >v<:\n", block.Header.Hash)
	fmt.Printf("\tHeight: '%v'\n", block.Header.Number)
	fmt.Printf("\tParent: '%x'\n\n", block.Header.Parent)
	fmt.Printf("\tNonce: '%v'\n", block.Header.Nonce)
	fmt.Printf("\tCreated: '%v'\n", block.Header.Time)
	fmt.Printf("\tTxs: '%x'\n", block.Header.Txs)

	fmt.Printf("\tAttempt: '%v'\n", attempt)
	fmt.Printf("\tSpent Time: %s\n\n", time.Since(start))

	return nil
}

func generateNonce() uint32 {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Uint32()
}

// var miningCtx context.Context
// var stopCurrentMining context.CancelFunc

// ticker := time.NewTicker(time.Second * intervalSeconds)
// defer ticker.Stop()

// for {
// 	select {
// 	case <-ticker.C:
// 		go func() {
// 			if len(n.pendingTXs) > 0 {

// 				miningCtx, stopCurrentMining = context.WithCancel(ctx)
// 				err := n.minePendingTXs(miningCtx)
// 				if err != nil {
// 					fmt.Printf("ERROR: %s\n", err)
// 				}

// 			}
// 		}()

// 	case <-ctx.Done():
// 		return nil
// 	}
// }
// }

// func (n *Node) minePendingTXs(ctx context.Context) error {
// 	blockToMine := NewPendingBlock(
// 		n.state.LatestBlockHash(),
// 		n.state.NextBlockNumber(),
// 		n.info.Account,
// 		n.getPendingTXsAsArray(),
// 	)

// 	minedBlock, err := Mine(ctx, blockToMine)
// 	if err != nil {
// 		return err
// 	}

// 	n.removeMinedPendingTXs(minedBlock)

// 	_, err = n.state.AddBlock(minedBlock)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// ticker := time.NewTicker(time.Second * msgTickSecs)
// 					defer ticker.Stop()

// 					done := make(chan bool)
// 					go func() {
// 						p.ProcessMessage(msg)
// 						done <- true
// 					}()

// 				LOOP:
// 					for {
// 						select {
// 						case <-done:
// 							break LOOP
// 						case <-ticker.C:
// 							if err := retry.Do(3, 300*time.Millisecond, func() error {
// 								_, err := p.ChangeMessageVisibilityWithContext(context.Background(), &sqs.ChangeMessageVisibilityInput{
// 									QueueUrl:          p.QueueUrl,
// 									ReceiptHandle:     msg.ReceiptHandle,
// 									VisibilityTimeout: aws.Int64(msgDelaySecs),
// 								})
// 								return err
// 							}); err != nil {
// 								logrus.WithField("messageId", msg.MessageId).Errorln("fail to delay task for time ticker:", err)
// 							}
// 						}
// 					}
