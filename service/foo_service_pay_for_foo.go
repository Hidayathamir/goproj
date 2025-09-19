package service

import (
	"context"
	"fmt"
)

func (s *FooServiceImpl) PayForFoo(ctx context.Context, req PayForFooReq) (PayForFooRes, error) {
	_, err := s.BackupFoo(ctx, BackupFooReq{ID: req.ID})
	if err != nil {
		return PayForFooRes{}, fmt.Errorf("s.BackupFoo: %w", err)
	}
	payRes, err := s.paymentService.Pay(ctx, PayReq{
		Username: req.Username,
		Password: req.Password,
		Amount:   req.Amount,
		Currency: req.Currency,
	})
	if err != nil {
		return PayForFooRes{}, fmt.Errorf("s.paymentService.Pay: %w", err)
	}
	return PayForFooRes{TransactionId: payRes.TransactionID}, nil
}
