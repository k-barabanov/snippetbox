package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct { // injecting dependencies into handlers (chapter 3.3)
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address. Usage: -addr=\":9999\"\n")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // custom logger 1/2

	app := &application{
		logger: logger,
	}

	fmt.Printf("starting server on http://localhost%s\n", *addr) // old version just to print port link in terminal
	logger.Info("starting server", "addr", *addr)                // custom logger 2/2

	err := http.ListenAndServe(*addr, app.routes())

	logger.Error(err.Error())
	os.Exit(1)
}
