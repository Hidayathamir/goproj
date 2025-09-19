package service

import (
	"context"

	"github.com/Hidayathamir/goproj/client"
	"github.com/Hidayathamir/goproj/repository"
)

type BarService interface {
	DoSomething(ctx context.Context, req DoSomethingReq) (DoSomethingRes, error)
	DoAnotherThing(ctx context.Context, req DoAnotherThingReq) (DoAnotherThingRes, error)
	GetAllData(ctx context.Context, req GetAllDataReq) (GetAllDataRes, error)
	NotifyAll(ctx context.Context, req NotifyAllReq) (NotifyAllRes, error)
}

// compile-time check
var _ BarService = (*BarServiceImpl)(nil)

type BarServiceImpl struct {
	// repository
	barRepository repository.BarRepository

	// client
	authClient  client.AuthClient
	slackClient client.SlackClient

	// service
}

func NewBarService(
	// repository
	barRepository repository.BarRepository,

	// client
	authClient client.AuthClient,
	slackClient client.SlackClient,

	// service
) *BarServiceImpl {
	return &BarServiceImpl{
		// repository
		barRepository: barRepository,

		// client
		authClient:  authClient,
		slackClient: slackClient,

		// service
	}
}
