package service

import (
	"context"

	"github.com/Hidayathamir/goproj/client"
	"github.com/Hidayathamir/goproj/repository"
)

type FooService interface {
	GetFoo(ctx context.Context, req GetFooReq) (GetFooRes, error)
	SaveFoo(ctx context.Context, req SaveFooReq) (SaveFooRes, error)
	BackupFoo(ctx context.Context, req BackupFooReq) (BackupFooRes, error)
	PayForFoo(ctx context.Context, req PayForFooReq) (PayForFooRes, error)
}

// compile-time check
var _ FooService = (*FooServiceImpl)(nil)

type FooServiceImpl struct {
	// repository
	fooRepository repository.FooRepository

	// client
	s3Client client.S3Client

	// service
	paymentService PaymentService
}

func NewFooService(
	// repository
	fooRepository repository.FooRepository,

	// client
	s3Client client.S3Client,

	// service
	paymentService PaymentService,
) *FooServiceImpl {
	return &FooServiceImpl{
		// repository
		fooRepository: fooRepository,

		// client
		s3Client: s3Client,

		// service
		paymentService: paymentService,
	}
}
