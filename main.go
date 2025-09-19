package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/Hidayathamir/goproj/cmd"
)

func main() {
	setupSlog()

	err := cmd.RunServer(context.Background())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func setupSlog() {
	var Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	slog.SetDefault(Logger)
}
