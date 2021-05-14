package main

import (
	"context"
	v1 "github.com/ethan510010/Week_04/api/book/v1"
	"github.com/ethan510010/Week_04/internal/service"
	rpc "github.com/ethan510010/Week_04/internal/pkg/grpc"
)

const address = ":8080"

func main() {
	bookUc := InitBookUsecase()
	service := service.NewUserService(bookUc)

	s := rpc.NewServer(address)
	v1.RegisterBookServiceServer(s.Server, service)

	s.Start(context.Background())
}
