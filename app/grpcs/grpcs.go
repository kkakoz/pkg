package grpcs

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	lis    net.Listener
	server *grpc.Server
}

func NewGrpcServer(addr string, server *grpc.Server) *GrpcServer {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return &GrpcServer{lis: lis, server: server}
}

func (g *GrpcServer) Start(ctx context.Context) error {

	if err := g.server.Serve(g.lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func (g GrpcServer) Stop(ctx context.Context) error {
	g.server.GracefulStop()
	return nil
}
