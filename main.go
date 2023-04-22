package main

import (
	"fmt"
	"net"

	"github.com/resource-aware-jds/common-go/proto"
	"github.com/resource-aware-jds/control-plane/handler"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	grpcHandler := handler.NewGRPCHandler()

	// GRPC
	lis, err := net.Listen("tcp", fmt.Sprint(":", 3000))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterControlPlaneServer(s, grpcHandler)

	logrus.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
