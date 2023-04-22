package handler

import "github.com/resource-aware-jds/common-go/proto"

type grpcHandler struct {
	proto.UnimplementedControlPlaneServer
}

func NewGRPCHandler() proto.ControlPlaneServer {
	return &grpcHandler{}
}
