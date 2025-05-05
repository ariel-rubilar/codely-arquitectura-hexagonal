package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/plataform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run(context.Context) error
}

type api struct {
	httpAddr string
	engine   *gin.Engine
}

func New(ctx context.Context, host string, port int) (context.Context, Server) {
	api := &api{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.Default(),
	}

	api.registerRoutes()

	return serverContext(ctx), api
}

func (a *api) Run(ctx context.Context) error {
	fmt.Println("Starting server on", a.httpAddr)

	srv := &http.Server{
		Addr:    a.httpAddr,
		Handler: a.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(ctx)
}

func (a *api) registerRoutes() {
	a.engine.GET("/health", health.CheckHandler())
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
