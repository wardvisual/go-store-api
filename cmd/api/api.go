package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wardvisual/go-store-api/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// User Resource
	userRepository := user.NewRepository(s.db)
	userrHandler := user.NewHandler(userRepository)
	userrHandler.RegisterRoutes(subrouter)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Server is running on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
