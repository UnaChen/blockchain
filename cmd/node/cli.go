package main

import (
	"blockchain/node"
	"blockchain/pb"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/spf13/cobra"
)

func runCmd() *cobra.Command {
	var port uint

	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Launch the node and its grpc server.",
		Run: func(cmd *cobra.Command, args []string) {

			n, err := node.New(port)
			if err != nil {
				logrus.Fatalln(port)
			}

			err = n.Run(context.Background())
			if err != nil {
				logrus.Fatalln(err)
			}
		},
	}
	runCmd.Flags().UintVarP(&port, "port", "p", 5566, "grpc sever port")
	return runCmd
}

func addTxCmd() *cobra.Command {
	var (
		addr  string
		from  string
		to    string
		value uint64
	)

	var cmd = &cobra.Command{
		Use:   "addtx",
		Short: "Launch the node and its grpc server.",
		Run: func(cmd *cobra.Command, args []string) {

			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				logrus.Fatalln(err)
			}
			defer conn.Close()

			client := pb.NewNodeClient(conn)
			resp, err := client.TxAdd(context.Background(), &pb.TxAddRequest{
				From:  from,
				To:    to,
				Value: value,
			})
			if err != nil {
				logrus.Fatalln(err)
			}

			output, _ := json.Marshal(resp)
			fmt.Println(string(output))
		},
	}

	cmd.Flags().StringVarP(&addr, "server", "s", "127.0.0.1:5566", "node grpc sever address")
	cmd.Flags().StringVarP(&from, "from", "f", genRandomString(), "from")
	cmd.Flags().StringVarP(&to, "to", "t", genRandomString(), "to")
	cmd.Flags().Uint64VarP(&value, "value", "v", 0, "value")
	return cmd
}

func listCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Short: "list info",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	cmd.AddCommand(balanceListCmd())
	cmd.AddCommand(blockListCmd())
	cmd.AddCommand(nodeStatusCmd())

	return cmd
}

func balanceListCmd() *cobra.Command {
	var addr string
	var cmd = &cobra.Command{
		Use:   "balance",
		Short: "List blances.",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				logrus.Fatalln(err)
			}
			defer conn.Close()

			client := pb.NewNodeClient(conn)
			resp, err := client.BalanceList(context.Background(), &pb.BalanceListRequest{})
			if err != nil {
				logrus.Fatalln(err)
			}

			output, _ := json.Marshal(resp)
			fmt.Println(string(output))
		},
	}

	cmd.Flags().StringVarP(&addr, "server", "s", "127.0.0.1:5566", "node grpc sever address")
	return cmd
}

func blockListCmd() *cobra.Command {
	var addr string
	var cmd = &cobra.Command{
		Use:   "block",
		Short: "List blocks.",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				logrus.Fatalln(err)
			}
			defer conn.Close()

			client := pb.NewNodeClient(conn)
			resp, err := client.BlockList(context.Background(), &pb.BlockListRequest{})
			if err != nil {
				logrus.Fatalln(err)
			}

			output, _ := json.Marshal(resp)
			fmt.Println(string(output))
		},
	}

	cmd.Flags().StringVarP(&addr, "server", "s", "127.0.0.1:5566", "node grpc sever address")
	return cmd
}

func nodeStatusCmd() *cobra.Command {
	var addr string
	var cmd = &cobra.Command{
		Use:   "status",
		Short: "Get node status.",
		Run: func(cmd *cobra.Command, args []string) {

			conn, err := grpc.Dial(addr, grpc.WithInsecure())
			if err != nil {
				logrus.Fatalln(err)
			}
			defer conn.Close()

			client := pb.NewNodeClient(conn)
			resp, err := client.NodeStatus(context.Background(), &pb.NodeStatusRequest{})
			if err != nil {
				logrus.Fatalln(err)
			}

			output, _ := json.Marshal(resp)
			fmt.Println(string(output))
		},
	}

	cmd.Flags().StringVarP(&addr, "server", "s", "127.0.0.1:5566", "node grpc sever address")
	return cmd
}

func genRandomString() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 32
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
