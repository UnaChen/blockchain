package main

import (
	"blockchain/pb"
	"flag"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int64("p", 5566, "grpc sever port")
	flag.Parse()

	s := grpc.NewServer()
	pb.RegisterNodeServer(s, &server{})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.WithField("netListen", listen).Errorln(err)
		panic(err)
	}
	defer listen.Close()
	defer func() {
		fmt.Println("defer")
	}()

	i := 0
	for {
		i++
		if i == 100 {
			return
		}
	}

	logrus.Infof("node grpc start on '%s' ...", listen.Addr())

	if err := s.Serve(listen); err != nil {
		logrus.WithField("serve", s).Errorln(err)
		panic(err)
	}

}
