package main

import (
	"blockchain/pb"
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	endpoint := flag.String("gprc", "localhost:5566", "grpc sever endpoint")
	port := flag.Int64("p", 7788, "grpc gateway port")
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterNodeHandlerFromEndpoint(ctx, mux, *endpoint, opts); err != nil {
		panic(err)
	}

	logrus.Infof("connect to node grpc sever on '%s' ...", *endpoint)
	logrus.Infof("run node grpc gateway on port '%d' ...", *port)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)

}
