package node

import (
	"blockchain/db"
	"blockchain/miner"
	"blockchain/pb"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Node struct {
	port       int
	state      *db.State
	pendingTXs map[string]*db.TX
	*miner.Miner
}

func NewNode(port int) (*Node, error) {

	gensisBlock, err := db.NewGensisBlock()
	if err != nil {
		return nil, err
	}

	return &Node{
		port:       port,
		state:      db.NewState(*gensisBlock),
		pendingTXs: make(map[string]*db.TX),
		Miner:      &miner.Miner{},
	}, nil

}

func (n *Node) Run(ctx context.Context) error {

	fmt.Println("Blockchain state:")
	fmt.Printf("	- height: %d\n", n.state.LatestBlockHeader.Number)
	fmt.Printf("	- hash: %x\n", n.state.LatestBlockHeader.Hash)

	s := grpc.NewServer()
	pb.RegisterNodeServer(s, n)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", n.port))
	if err != nil {
		return err
	}
	defer listen.Close()

	logrus.Infof("node grpc start on '%s' ...", listen.Addr())

	go n.mine(ctx)

	return s.Serve(listen)
}

func (n *Node) TxAdd(ctx context.Context, req *pb.TxAddRequest) (*pb.TxAddResponse, error) {

	if req.From == "" || req.To == "" {
		return nil, fmt.Errorf("from/to is empty")
	}

	from := common.HexToAddress(req.From)

	tx := &db.TX{
		From:  from,
		To:    common.HexToAddress(req.To),
		Value: uint(req.Value),
		Nonce: n.state.Account2Nonce[from] + 1,
	}

	if err := db.NewTX(tx); err != nil {
		return nil, err
	}

	n.pendingTXs[tx.Hash] = tx

	msg, _ := json.Marshal(tx)

	return &pb.TxAddResponse{
		Status:  "success",
		Message: string(msg),
	}, nil
}

func (n *Node) BalanceList(ctx context.Context, request *pb.BalanceListRequest) (*pb.BalanceListResponse, error) {
	balances := n.state.GetBalances()

	output := map[string]uint64{}
	for a, v := range balances {
		output[a.String()] = uint64(v)
	}

	return &pb.BalanceListResponse{
		Balances: output,
	}, nil
}

func (n *Node) BlockList(ctx context.Context, request *pb.BlockListRequest) (*pb.BlockListResponse, error) {
	blocks := n.state.GetBlocks()
	data, _ := json.Marshal(blocks)

	resp := &pb.BlockListResponse{}
	err := json.Unmarshal(data, &resp.Blcoks)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *Node) NodeStatus(ctx context.Context, request *pb.NodeStatusRequest) (*pb.NodeStatusResponse, error) {
	txs := n.pendingTXs
	data, _ := json.Marshal(txs)

	resp := &pb.NodeStatusResponse{}
	err := json.Unmarshal(data, &resp.PendingTxs)
	if err != nil {
		return nil, err
	}
	latestBlockerHeader := n.state.LatestBlockHeader
	resp.BlockHeight = latestBlockerHeader.Number
	resp.BlockLatestHash = string(latestBlockerHeader.Hash[:])

	return resp, nil
}

func (n *Node) mine(ctx context.Context) {

	ticker := time.NewTicker(time.Second * miner.IntervalSeconds)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if len(n.pendingTXs) > 0 {

				err := n.minePendingTXs(ctx)
				if err != nil {
					logrus.Errorln(err)
				}

			}

		case <-ctx.Done():
			return
		}
	}
}
func (n *Node) minePendingTXs(ctx context.Context) error {
	pendingBlock := &db.Block{
		Header: db.BlockHeader{
			Parent: n.state.LatestBlockHeader.Hash,
			Number: n.state.LatestBlockHeader.Number + 1,
		},
	}

	for hash, tx := range n.pendingTXs {
		pendingBlock.Header.TXs = append(pendingBlock.Header.TXs, hash)
		pendingBlock.TXs = append(pendingBlock.TXs, *tx)
	}

	err := db.NewBlock(pendingBlock)
	if err != nil {
		return errors.Wrap(err, "fail to create a pending block")
	}

	err = n.Mine(ctx, pendingBlock)
	if err != nil {
		return errors.Wrap(err, "fail to mine a block")
	}

	err = n.state.AddBlock(*pendingBlock)
	if err != nil {
		n.removeMinedPendingTXs(*pendingBlock)
		return errors.Wrap(err, "fail to apply a new block")
	}

	n.removeMinedPendingTXs(*pendingBlock)
	return nil
}

func (n *Node) removeMinedPendingTXs(block db.Block) {
	if len(block.TXs) > 0 && len(n.pendingTXs) > 0 {
		fmt.Println("Updating in-memory Pending TXs Pool")
	}

	for _, hash := range block.Header.TXs {
		if _, exists := n.pendingTXs[hash]; exists {
			fmt.Printf("\t-archiving mined TX: %s\n", hash)
			delete(n.pendingTXs, hash)
		}
	}
}
