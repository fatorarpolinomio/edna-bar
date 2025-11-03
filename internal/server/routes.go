package server

import (
	"edna/internal/services/fornecedor"
	"encoding/json"
	"log"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "edna/docs"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	fornecedorHandler := fornecedor.NewHandler(s.fornecedorStore)
	// Register routes
	mux.HandleFunc("/", s.trailingSlashHandler)
	mux.HandleFunc("/health", s.healthHandler)
	mux.Handle("/swagger/", httpSwagger.Handler())

	fornecedorHandler.RegisterRoutes(mux)

	// Wrap the mux with CORS middleware
	return s.logMiddleware(s.corsMiddleware(mux))
}

func (s *Server) trailingSlashHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	resp := map[string]string{"message": "Unmatched path, please check your url path and try again."}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

// Check health of the system
// @summary Check health of the system
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(s.db.Health())
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
