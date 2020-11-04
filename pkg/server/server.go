package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
	router chi.Router
	server *http.Server
}

func New(logger *logrus.Logger) (*Server, error) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	s := &Server{
		logger: logger,
		router: r,
	}

	return s, nil
}

func (s *Server) ListenAndServe() error {
	s.server = &http.Server{
		Addr:    net.JoinHostPort("127.0.0.1", "8080"),
		Handler: s.router,
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("Could not listen on %s: %v", s.server.Addr, err)
	}

	go func() {
		if err = s.server.Serve(listener); err != nil {
			s.logger.Fatalf("API Listen error `%v` address `%v`", err, s.server.Addr)
		}
	}()
	s.logger.Infof("API Listening `%v`", s.server.Addr)

	return nil
}

// Router returns the router
func (s *Server) Router() chi.Router {
	return s.router
}
