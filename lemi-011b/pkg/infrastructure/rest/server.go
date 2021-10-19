package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"
)

// Server TODO
type Server struct {
	mux      *http.ServeMux
	acquirer acquisition.Service
}

// NewServer TODO
func NewServer(
	acquisitionService acquisition.Service,
) (*Server, error) {
	server := &Server{
		mux:      http.NewServeMux(),
		acquirer: acquisitionService,
	}

	server.mux.HandleFunc("/datum", server.AcquireDatum)

	return server, nil
}

// ServeHTTP TODO
func (server *Server) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	server.mux.ServeHTTP(w, r)
}

// AcquireDatum TODO
func (server *Server) AcquireDatum(
	w http.ResponseWriter,
	r *http.Request,
) {
	datum := acquisition.Datum{}

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&datum)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = server.acquirer.AcquireDatum(r.Context(), datum)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
