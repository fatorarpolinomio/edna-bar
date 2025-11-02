package fornecedor

import (
	"context"
	"edna/internal/model"
	"edna/internal/util"
	"encoding/json"
	"net/http"
	"time"
)

var (
	requestTimeout = 2 * time.Second
)

type Handler struct {
	store FornecedorStore
}

type FornecedorStore interface {
	GetAll(ctx context.Context) ([]model.Fornecedor, error)
	Create(ctx context.Context, props *model.Fornecedor) error
}


func NewHandler(store FornecedorStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /fornecedores", h.getAllFornecedoresHandler)
	mux.HandleFunc("POST /fornecedores", h.createFornecedoresHandler)
}


func (h *Handler) getAllFornecedoresHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	fornecedores, err := h.store.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = util.WriteJSON(w, http.StatusOK, fornecedores)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) createFornecedoresHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), requestTimeout)
	defer cancel()

	if r.Body == nil {
		http.Error(w, "No body in the request", http.StatusBadRequest)
		return
	}

	var model model.Fornecedor
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.store.Create(ctx, &model)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.WriteJSON(w, http.StatusCreated, model)
}
