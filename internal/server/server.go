package server

import (
	"net/http"

	"github.com/micro/go-micro/util/log"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/orlovssky/gread/internal/routes"
	"github.com/orlovssky/gread/internal/secrets"
	"github.com/orlovssky/gread/internal/store"
	"github.com/orlovssky/gread/pkg/logger"
)

// Server type
type server struct {
	log     *logrus.Logger
	db      *gorm.DB
	secrets secrets.Secrets
}

// NewServer - Create Server
func NewServer() server {
	// Create a new server
	s := server{}

	// Inject logger
	s.log = logger.NewLogger()

	// Inject secrets
	var err error
	s.secrets, err = secrets.LoadSecrets()
	if err != nil {
		s.log.Error("error:", err)
	}

	// Inject database
	s.db = store.NewDatabase(s.secrets)
	return s
}

// StartServer - Load routes into server and
// starts HTTP server
func (s *server) StartServer() {
	log.Info("📡 Server Started. API Server is now listening on http://localhost:8080")
	routes := routes.Routes()
	log.Fatal(http.ListenAndServe(":8080", routes))
}
