package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/micro/go-micro/util/log"
	"github.com/orlovssky/gread/internal/logger"
	"github.com/orlovssky/gread/internal/routes"
	"github.com/sirupsen/logrus"

	"github.com/orlovssky/gread/internal/secrets"
	"github.com/orlovssky/gread/internal/store"
)

// Server type
type server struct {
	r   chi.Router
	log *logrus.Logger
}

// NewServer - Create Server
func NewServer() server {
	// Create a new server
	s := server{}

	// Inject logger
	s.log = logger.NewLogger()

	// Load secrets
	var err error
	_, err = secrets.LoadSecrets()
	if err != nil {
		s.log.Error("error:", err)
	}

	// Connect to DB
	store.ConnectToDB()

	return s
}

// StartServer - Load routes into server and
// starts HTTP server
func (s *server) StartServer() {
	log.Info("📡 Server Started. API Server is now listening on http://localhost:8081")
	s.r = routes.Routes()
	log.Fatal(http.ListenAndServe(":8081", s.r))
}

// ServeHTTP - Turns server into http server
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}
