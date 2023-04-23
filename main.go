package main

import (
	"context"
	"fmt"
	"github.com/resource-aware-jds/common-go/logger"
	"github.com/resource-aware-jds/control-plane/config"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	"github.com/resource-aware-jds/common-go/proto"
	"github.com/resource-aware-jds/control-plane/handler"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	appConfig config.Config
)

func init() {
	appConfig = config.Load()
	logger.InitLogger(logger.Config{
		Env: appConfig.Env,
	})

}

func main() {
	grpcHandler := handler.NewGRPCHandler()

	go func() {
		// GRPC
		lis, err := net.Listen("tcp", fmt.Sprint(":", appConfig.GRPC_SERVER_PORT))
		if err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		proto.RegisterControlPlaneServer(s, grpcHandler)

		logrus.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			logrus.Fatalf("failed to serve: %v", err)
		}
	}()

	exampleNodePath := "127.0.0.1:3001"
	ctx := context.Background()
	grpcConn, err := grpc.Dial(exampleNodePath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic("Can't connect gRPC Server: ", err)
	}
	defer grpcConn.Close()

	computeNodeClient := proto.NewComputeNodeClient(grpcConn)
	_, err = computeNodeClient.SendJob(ctx, &proto.Job{JobID: 1, DockerImage: "ghcr.io/resource-aware-jds/example"})
	fmt.Println(err)
}
