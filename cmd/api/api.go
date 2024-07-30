package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	// subrouter = router.NewRoute().Subrouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Server is running on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
