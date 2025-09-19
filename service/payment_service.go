package service

import (
	"context"

	"github.com/Hidayathamir/goproj/client"
)

type PaymentService interface {
	Pay(ctx context.Context, req PayReq) (PayRes, error)
	Refund(ctx context.Context, req RefundReq) (RefundRes, error)
}

// compile-time check
var _ PaymentService = (*PaymentServiceImpl)(nil)

type PaymentServiceImpl struct {
	// repository

	// client
	paymentClient client.PaymentClient
	authClient    client.AuthClient

	// service
}

func NewPaymentService(
	// repository

	// client
	paymentClient client.PaymentClient,
	authClient client.AuthClient,

	// service
) *PaymentServiceImpl {
	return &PaymentServiceImpl{
		// repository

		// client
		paymentClient: paymentClient,
		authClient:    authClient,

		// service
	}
}
