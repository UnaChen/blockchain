package main

import (
	"blockchain/node"
	"blockchain/pb"
	"context"
	"fmt"
)

type server struct {
	*node.Node
	// *gorm.DB
	// cloudstorage.Client
	// vt *virustotal.Client
	// // *s3manager.Downloader
	// // *s3manager.Uploader
	// *session.Session
	// S3Bucket string
}

// func (s *server) init(){
// 	s.
// }

func (s *server) TXAdd(ctx context.Context, request *pb.TXAddRequest) (*pb.TXAddResponse, error) {
	return &pb.TXAddResponse{}, nil
}

func (s *server) BalancesList(ctx context.Context, request *pb.BalancesListRequest) (*pb.BalancesListResponse, error) {
	fmt.Println("BalancesList	")
	return &pb.BalancesListResponse{}, nil
}

func (s *server) NodeStatus(ctx context.Context, request *pb.NodeStatusRequest) (*pb.NodeStatusResponse, error) {
	return &pb.NodeStatusResponse{}, nil
}
