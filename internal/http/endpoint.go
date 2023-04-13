package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) setupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/alive", func(w http.ResponseWriter, r *http.Request) {
		s.cfg.L.Info("alive endpoint")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<h1>Server is working</h1>"))

		return
	})

	return r
}
