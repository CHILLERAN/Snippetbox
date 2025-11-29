package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Application struct{
	logger *slog.Logger
}

func main() {
	address := flag.String("addr", ":4000", "HTTP Network Address")

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})
	logger := slog.New(loggerHandler)

	app := &Application{
		logger: logger,
	}

	flag.Parse()

	app.logger.Info("Starting server", "Address", fmt.Sprintf("http://localhost%v", *address))

	err := http.ListenAndServe(*address, app.routes())

	app.logger.Error(err.Error())
	os.Exit(1)
}