package main

import (
	"finance-tracker-system/transaction-service/config"
	"finance-tracker-system/transaction-service/internal/handler"
	"finance-tracker-system/transaction-service/internal/repository"
	"finance-tracker-system/transaction-service/internal/server"
	"finance-tracker-system/transaction-service/internal/service"
	"fmt"
)

func main() {
	cfg := config.NewConfig()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB_user, cfg.DB_password, cfg.DB_address, cfg.DB_port, cfg.DB_name)
	pool := repository.NewPostgresPool(connString)

	repo := repository.NewTransactionRepository(pool)
	service := service.NewTransactionService(repo)
	handler := handler.NewTransactionHandler(service)

	server.StartGRPCServer(handler)
}
