package cmd

import (
	"github.com/Hidayathamir/goproj/client"
	"github.com/Hidayathamir/goproj/controller"
	"github.com/Hidayathamir/goproj/repository"
	"github.com/Hidayathamir/goproj/service"
)

type dependencyInjection struct {
	fooController *controller.FooController
	barController *controller.BarController
}

func wireDependencyInjection() *dependencyInjection {
	// repository
	fooRepository := repository.NewFooRepository()
	barRepository := repository.NewBarRepository()

	// client
	slackClient := client.NewSlackClient()
	s3Client := client.NewS3Client()
	paymentClient := client.NewPaymentClient()
	authClient := client.NewAuthClient()

	// service
	paymentService := service.NewPaymentService(paymentClient, authClient)
	fooService := service.NewFooService(fooRepository, s3Client, paymentService)
	barService := service.NewBarService(barRepository, authClient, slackClient)

	// controller
	fooController := controller.NewFooController(fooService)
	barController := controller.NewBarController(barService)

	return &dependencyInjection{
		fooController: fooController,
		barController: barController,
	}
}
