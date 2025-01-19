package web

import (
	"context"
	"fmt"
	"go-api-arch-clean/adapter/controller/gin/router"
	"go-api-arch-clean/pkg/logger"
	"net/http"

	"gorm.io/gorm"
)

type GinWebServer struct {
	server *http.Server
}

func (g *GinWebServer) Start() error {
	return g.server.ListenAndServe()

}

func (g *GinWebServer) Shutdown(ctx context.Context) error {
	return g.server.Shutdown(ctx)
}

func NewGinServer(host, port string, corsAllowOrigins []string, db *gorm.DB) (Server, error) {
	router, err := router.NewGinRouter(db, corsAllowOrigins)
	if err != nil {
		logger.Error(err.Error(), "host", host, "port", port)
		return nil, err
	}
	return &GinWebServer{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", host, port),
			Handler: router,
		},
	}, err
}
