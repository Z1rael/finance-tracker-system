package main

import (
	"finance-tracker-system/transaction-service/internal/handler"
	"finance-tracker-system/transaction-service/internal/repository"
	"finance-tracker-system/transaction-service/internal/server"
	"finance-tracker-system/transaction-service/internal/service"
)

func main() {
	repo := repository.NewInMemoryTransactionRepository()
	service := service.NewTransactionService(repo)
	handler := handler.NewTransactionHandler(service)

	server.StartGRPCServer(handler)
}
