package handler

import (
	"context"
	"finance-tracker-system/transaction-service/internal/service"
	pb "finance-tracker-system/transaction-service/proto/gen/go/transaction"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TransactionHandler struct {
	pb.UnimplementedTransactionServiceServer
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (handler *TransactionHandler) CreateTransaction(
	ctx context.Context,
	req *pb.CreateTransactionRequest,
) (*pb.CreateTransactionResponse, error) {
	t, err := handler.service.CreateTransaction(
		req.AccountId,
		req.Amount,
		req.Description,
		int32(req.Category),
		int32(req.TransactionType),
		req.Timestamp.AsTime(),
	)

	if err != nil {
		return nil, err
	}

	return &pb.CreateTransactionResponse{
		Transaction: &pb.Transaction{
			Id:              t.ID,
			AccountId:       t.AccountID,
			Amount:          t.Amount,
			Description:     t.Description,
			Category:        req.Category,
			TransactionType: req.TransactionType,
			Timestamp:       timestamppb.New(t.Timestamp),
		},
	}, nil
}

func (handler *TransactionHandler) ListTransactions(
	ctx context.Context,
	req *pb.ListTransactionsRequest,
) (*pb.ListTransactionsResponse, error) {
	list, err := handler.service.ListTransactions(req.AccountId)
	if err != nil {
		return nil, err
	}

	var result []*pb.Transaction
	for _, t := range list {
		result = append(result, &pb.Transaction{
			Id:              t.ID,
			AccountId:       t.AccountID,
			Amount:          t.Amount,
			Description:     t.Description,
			Category:        pb.TransactionCategory(t.Category),
			TransactionType: pb.TransactionType(t.Type),
			Timestamp:       timestamppb.New(t.Timestamp),
		})
	}

	return &pb.ListTransactionsResponse{
		Transactions: result,
	}, nil
}

func (handler *TransactionHandler) DeleteTransaction(
	ctx context.Context,
	req *pb.DeleteTransactionRequest,
) (*emptypb.Empty, error) {
	err := handler.service.DeleteTransaction(req.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
