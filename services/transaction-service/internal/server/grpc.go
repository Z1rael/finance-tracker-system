package server

import (
	"finance-tracker-system/transaction-service/internal/handler"
	pb "finance-tracker-system/transaction-service/proto/gen/go/transaction"
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPCServer(handler *handler.TransactionHandler) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTransactionServiceServer(grpcServer, handler)

	log.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
