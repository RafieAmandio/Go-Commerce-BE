package server

import (
	"mooi/config"
	"mooi/library/logger"
	"mooi/storage/db"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var once sync.Once
var server Server

type Server struct {
	Router   *mux.Router
	Database db.PostgreSQLDB
}

func NewServer() Server {
	once.Do(func() {
		server = Server{}
		server.Router = mux.NewRouter()
		server.routes()
	})
	return server
}

func (s *Server) SetDatabase(database db.PostgreSQLDB) {
	s.Database = database
}

func (s *Server) StartHttpServer() error {

	appPort := config.Get().AppPort
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         ":" + appPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Log.Info("Server running on port: ", appPort)

	err := srv.ListenAndServe()
	return err
}
