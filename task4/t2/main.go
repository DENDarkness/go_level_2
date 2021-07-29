package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	httpServer *http.Server
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run() error {

	app.httpServer = &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 1*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)

}

func main() {

	app := NewApp()

	if err := app.Run(); err != nil {
		log.Fatalf("%s", err.Error())
	}

}
