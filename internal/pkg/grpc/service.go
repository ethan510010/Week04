package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*grpc.Server
	address string
}

func NewServer(address string) *Server {
	server := grpc.NewServer()
	return &Server{
		server,
		address,
	}
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("grpc server start: %s", s.address)

	go func() {
		<- ctx.Done()
		s.GracefulStop()
		fmt.Println("grpc server stop gracefully")
	}()

	return s.Serve(lis)
}