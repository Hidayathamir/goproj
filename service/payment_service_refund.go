package service

import (
	"context"
	"errors"
)

func (s *PaymentServiceImpl) Refund(ctx context.Context, req RefundReq) (RefundRes, error) {
	token := s.authClient.Login(req.Username, req.Password)
	if !s.authClient.ValidateToken(token) {
		return RefundRes{}, errors.New("invalid token")
	}
	ok := s.paymentClient.Refund(req.TransactionID)
	return RefundRes{Success: ok}, nil
}
