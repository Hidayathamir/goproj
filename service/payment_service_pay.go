package service

import (
	"context"
	"errors"
)

func (s *PaymentServiceImpl) Pay(ctx context.Context, req PayReq) (PayRes, error) {
	token := s.authClient.Login(req.Username, req.Password)
	if !s.authClient.ValidateToken(token) {
		return PayRes{}, errors.New("invalid token")
	}
	transactionID := s.paymentClient.Charge(req.Amount, req.Currency)
	return PayRes{TransactionID: transactionID}, nil
}
