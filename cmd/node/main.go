package main

import (
	"blockchain/node"
	"context"
	"flag"

	"github.com/sirupsen/logrus"
)

func main() {
	port := flag.Int("p", 5566, "grpc sever port")
	flag.Parse()

	node, err := node.NewNode(*port)
	if err != nil {
		logrus.Fatalf("fail to create node: %s", err)
	}

	node.Run(context.Background())
}
