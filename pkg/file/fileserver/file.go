package fileserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/snowzach/gorestapi/server"
	"github.com/yurifrl/logapi"
)

var (
	fileName = "examples/log.txt"
)

// Server is the API web server
type Server struct {
	logger *logrus.Logger
	router chi.Router
	store  logapi.FileStore
	file   logapi.File
}

// Setup will setup the API listener
func Setup(logger *logrus.Logger, router chi.Router, store logapi.FileStore, file logapi.File) error {
	s := &Server{
		logger: logger,
		router: router,
		store:  store,
		file:   file,
	}

	// Base Functions
	s.router.Get("/files", s.Files())

	return nil
}

func (s *Server) Files() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.file.Sync(fileName)
		if err != nil {
			render.Render(w, r, server.ErrInvalidRequest(err))
			return
		}

		bs, err := s.store.GetAll()
		if err != nil {
			render.Render(w, r, server.ErrInvalidRequest(err))
			return
		}

		render.JSON(w, r, bs)
	}
}
