package cmd

import (
	"github.com/Hidayathamir/goproj/pkg/http/middleware"
	"github.com/gin-gonic/gin"
)

func registerRoutes(ginRouter *gin.RouterGroup, dependencyInjection *dependencyInjection) {
	foo := ginRouter.Group("/foo", middleware.DummyMiddleware())
	{
		foo.GET("/get", dependencyInjection.fooController.GetFoo)
		foo.POST("/save", dependencyInjection.fooController.SaveFoo)
		foo.POST("/backup", dependencyInjection.fooController.BackupFoo)
		foo.POST("/pay", dependencyInjection.fooController.PayForFoo)
	}

	bar := ginRouter.Group("/bar", middleware.DummyMiddleware())
	{
		bar.GET("/do-something", dependencyInjection.barController.DoSomething)
		bar.GET("/do-another-thing", dependencyInjection.barController.DoAnotherThing)
		bar.GET("/get-all", dependencyInjection.barController.GetAllData)
		bar.POST("/notify-all", dependencyInjection.barController.NotifyAll)
	}
}
