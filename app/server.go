package app

import (
  "fmt"
  "io"
  "database/sql"
  "net/http"
  "strings"
  "github.com/gorilla/mux"
)


type Server struct {
   db  *sql.DB
   router *mux.Router
}

func NewServer(msql *sql.DB, r *mux.Router) *Server{
	return &Server{db: msql,router: r}
}


func (s *Server) rootFuncHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        _, err := s.db.Query("show tables")
	if err != nil {
		w.WriteHeader(500)
		em := fmt.Sprintf("Error Connection to db: %s\n", err)
		msg := strings.NewReader(em)
		if _, err := io.Copy(w, msg); err != nil {
		   // do nothing 
		}
	}else{
		w.WriteHeader(200)
		msg := strings.NewReader("Root Handler: Successfully connected to mysql database")
		if _, err := io.Copy(w, msg); err != nil {
		   // do nothing 
		}
	}

    }
}

func (s *Server) Routes(){
    s.router.HandleFunc("/", s.rootFuncHandler())
}
