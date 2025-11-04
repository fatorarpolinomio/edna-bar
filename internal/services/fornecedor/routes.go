package fornecedor

import (
	"context"
	"edna/internal/model"
	"edna/internal/util"
	"encoding/json"
	"net/http"
)


type Handler struct {
	store FornecedorStore
}


type FornecedorStore interface {
	GetAll(ctx context.Context, filters model.FornecedorFilters) ([]model.Fornecedor, error)
	Create(ctx context.Context, props *model.Fornecedor) error
	GetByID(ctx context.Context, id int64) (*model.Fornecedor, error)
	Update(ctx context.Context, props *model.Fornecedor) error
	Delete(ctx context.Context, id int64) (*model.Fornecedor, error)
}


func NewHandler(store FornecedorStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /fornecedores", h.getAllFornecedoresHandler)
	mux.HandleFunc("POST /fornecedores", h.createFornecedoresHandler)
	mux.HandleFunc("GET /fornecedores/{id}", h.fetchFornecedorHandler)
	mux.HandleFunc("PUT /fornecedores/{id}", h.updateFornecedorHandler)
	mux.HandleFunc("DELETE /fornecedores/{id}", h.deleteFornecedorHandler)
}

// @Summary List Fornecedores
// @Tags Fornecedor
// @Produce json
// @Param nome query string false "Filter by nome (partial match)"
// @Param offset query int false "Pagination offset (default 0)"
// @Param limit query int false "Pagination limit (default 10)"
// @Param sort query string false "Sort order: asc or desc"
// @Success 200 {array} model.Fornecedor
// @Failure 500 {object} map[string]string
// @Router /fornecedores [get]
func (h *Handler) getAllFornecedoresHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	filters := model.NewFornecedorFilter(r.URL)
	fornecedores, err := h.store.GetAll(ctx, filters)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = util.WriteJSON(w, http.StatusOK, fornecedores)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Create Fornecedor
// @Tags Fornecedor
// @Accept json
// @Produce json
// @Param fornecedor body model.FornecedorPayload true "Fornecedor payload"
// @Success 201 {object} model.Fornecedor
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /fornecedores [post]
func (h *Handler) createFornecedoresHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	if r.Body == nil {
		util.ErrorJSON(w, "No body in the request", http.StatusBadRequest)
		return
	}

	var payload model.FornecedorPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := model.FromPayload(payload)
	err = h.store.Create(ctx, &model)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.WriteJSON(w, http.StatusCreated, model)
}

// @Summary Get Fornecedor by ID
// @Tags Fornecedor
// @Produce json
// @Param id path int true "Fornecedor ID"
// @Success 200 {object} model.Fornecedor
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /fornecedores/{id} [get]
func (h *Handler) fetchFornecedorHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	fornecedor, err := h.store.GetByID(ctx, id)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if fornecedor == nil {
		util.ErrorJSON(w, "Fornecedor not found.", http.StatusNotFound)
		return
	}

	if err = util.WriteJSON(w, http.StatusOK, fornecedor); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Update Fornecedor
// @Tags Fornecedor
// @Accept json
// @Produce json
// @Param id path int true "Fornecedor ID"
// @Param fornecedor body model.FornecedorPayload true "Fornecedor payload"
// @Success 200 {object} model.Fornecedor
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /fornecedores/{id} [put]
func (h *Handler) updateFornecedorHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload model.FornecedorPayload
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := model.FromPayload(payload)
	model.Id = id
	err = h.store.Update(ctx, &model)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.WriteJSON(w, http.StatusOK, model)
}

// @Summary Delete Fornecedor
// @Tags Fornecedor
// @Produce json
// @Param id path int true "Fornecedor ID"
// @Success 200 {object} model.Fornecedor
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Router /fornecedores/{id} [delete]
func (h *Handler) deleteFornecedorHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	model, err := h.store.Delete(ctx, id)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.WriteJSON(w, http.StatusOK, model)
}