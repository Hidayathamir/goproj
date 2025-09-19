package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func RunServer(ctx context.Context) error {
	dependencyInjection := wireDependencyInjection()

	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	registerRoutes(&ginEngine.RouterGroup, dependencyInjection)

	httpServer := &http.Server{Addr: ":8080", Handler: ginEngine.Handler()}

	go runHTTPServer(httpServer)

	slog.Info("listens for the interrupt signal from the OS")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err := gracefulShutdownHttpServer(httpServer)
	if err != nil {
		return fmt.Errorf("gracefulShutdownHttpServer: %w", err)
	}

	return nil
}

func runHTTPServer(httpServer *http.Server) {
	slog.Info("running http server", "address", httpServer.Addr)

	err := httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error(err.Error())
		return
	}
}

func gracefulShutdownHttpServer(httpServer *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("shutting down gracefully http server, inform http server it has 10 seconds to finish the request it is currently handling")
	err := httpServer.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("httpServer.Shutdown: %w", err)
	}

	slog.Info("shutting down gracefully http server done")
	return nil
}
