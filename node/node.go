package node

import (
	"blockchain/db"
	"blockchain/miner"
	"context"
	"fmt"

	"github.com/web3coach/the-blockchain-bar/database"
)

type Node struct {
	state      *db.State
	pendingTXs []db.Tx
	*miner.Miner
}

func (n *Node) Run(ctx context.Context, isSSLDisabled bool, sslEmail string) error {
	fmt.Println(fmt.Sprintf("Listening on: %s:%d", n.info.IP, n.info.Port))

	state, err := database.NewStateFromDisk(n.dataDir)
	if err != nil {
		return err
	}
	defer state.Close()

	n.state = state

	fmt.Println("Blockchain state:")
	fmt.Printf("	- height: %d\n", n.state.LatestBlock().Header.Number)
	fmt.Printf("	- hash: %s\n", n.state.LatestBlockHash().Hex())

	go n.sync(ctx)
	go n.mine(ctx)

	return n.serveHttp(ctx, isSSLDisabled, sslEmail)
}

// run
// go run mine
// mine

func (n *Node) AddPendingTX(tx database.Tx) {
	n.pendingTXs = append(n.pendingTXs)
}
