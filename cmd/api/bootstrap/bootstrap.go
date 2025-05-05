package bootstrap

import (
	"context"

	"github.com/ariel-rubilar/codely-arquitectura-hexagonal/internal/plataform/server"
)

const (
	host = "localhost"
	port = 8082
)

func Run() error {

	ctx, srv := server.New(context.Background(), host, port)

	return srv.Run(ctx)
}
