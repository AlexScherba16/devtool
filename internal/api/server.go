package api

import (
	"context"
	"devtool/config"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	ctx        context.Context
}

func NewServer(cfg *config.ServerConfig, handler http.Handler) *Server {
	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	ctx, _ := context.WithCancel(context.Background())
	return &Server{
		httpServer: s,
		ctx:        ctx,
	}
}

func (s *Server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error occurred on listen and serve: %v", err)
		}
	}()
	log.Println("Server is running: ", s.httpServer.Addr)
}

func (s *Server) Stop() error {
	return s.httpServer.Shutdown(s.ctx)
}
