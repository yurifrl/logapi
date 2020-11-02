package fileserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/snowzach/gorestapi/server"
	"github.com/yurifrl/logapi"
)

// Server is the API web server
type Server struct {
	logger *logrus.Logger
	router chi.Router
	store  logapi.FileStore
}

// Setup will setup the API listener
func Setup(logger *logrus.Logger, router chi.Router, store logapi.FileStore) error {
	s := &Server{
		logger: logger,
		router: router,
		store:  store,
	}

	// Base Functions
	s.router.Get("/metrics", s.Metrics())

	return nil
}

func (s *Server) Metrics() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bs, err := s.store.GetAll()
		if err != nil {
			render.Render(w, r, server.ErrInvalidRequest(err))
			return
		}

		render.JSON(w, r, bs)
	}
}
