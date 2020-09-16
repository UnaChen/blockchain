package node

import (
	"blockchain/db"
	"blockchain/miner"
	"blockchain/pb"
	"context"
	"encoding/json"
	"fmt"
	"net"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Node struct {
	port       int
	state      *db.State
	pendingTXs []db.TX
	*miner.Miner
}

func NewNode(port int) (*Node, error) {

	gensisBlock, err := db.NewGensisBlock()
	if err != nil {
		return nil, err
	}

	return &Node{
		port:  port,
		state: db.NewState(*gensisBlock),
		Miner: &miner.Miner{},
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

	// go n.mine(ctx)

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
		Nonce: n.state.GetNextAccountNonce(from),
	}

	if err := db.NewTX(tx); err != nil {
		return nil, err
	}

	n.pendingTXs = append(n.pendingTXs, *tx)

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

	// output := []*pb.Blcok{}
	// for _, b := range n.state.GetBlocks() {
	// 	output = append(output, &pb.Blcok{
	// 	Header:pb.BlockHeader{
	// 		Hash :
	// ParentHash string   `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	// Number     int64    `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// TxHashes   []string `protobuf:"bytes,4,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
	// Nonce      int64    `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Timestamp  int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 	}
	// Hash :b.hash,

	// ParentHash string   `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	// Number     int64    `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// TxHashes   []string `protobuf:"bytes,4,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
	// Nonce      int64    `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// // Timestamp  int64    `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 	})
	// }

	// return &pb.BlockListResponse{
	// 	Blcoks: output,
	// }, nil
	return nil, nil
}

func (s *Node) NodeStatus(ctx context.Context, request *pb.NodeStatusRequest) (*pb.NodeStatusResponse, error) {
	return &pb.NodeStatusResponse{}, nil
}
