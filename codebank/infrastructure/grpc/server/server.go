package server

import (
	"log"
	"net"

	"github.com/Valdeci97/fullcycle-event/infrastructure/grpc/pb"
	"github.com/Valdeci97/fullcycle-event/infrastructure/grpc/service"
	"github.com/Valdeci97/fullcycle-event/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	ProcessTransactionUseCase usecase.UseCaseTransaction
}

func NewGRPCServer() GRPCServer {
	return GRPCServer{}
}

func (s GRPCServer) Server() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatal("could not listen tcp port")
	}
	transactionService := service.NewTransactionService()
	transactionService.ProcessTransactionUseCase = s.ProcessTransactionUseCase
	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterPaymentServiceServer(server, transactionService)
	server.Serve(lis)
}
