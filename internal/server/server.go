package server

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	server *http.Server
	logger *zap.Logger
}

func NewServer(port string, parrentLoger *zap.Logger, router http.Handler) *Server {

	log := parrentLoger.With(zap.Namespace("http server"))

	serv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return &Server{
		server: &serv,
		logger: log,
	}
}

func (s *Server) Run() error {
	log := s.logger.With(zap.Namespace("http run")).With(zap.String("addr", s.server.Addr))

	err := s.server.ListenAndServe()

	log.Info("listen on http completed", zap.Error(err))

	if !errors.Is(err, http.ErrServerClosed) {
		log.Error("listen http error", zap.Error(err))
	}

	return err
}

func (s *Server) Shutdown() error {
	return s.server.Shutdown(context.TODO())
}
