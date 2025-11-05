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
	GetAll(ctx context.Context, filter util.Filter) ([]model.Fornecedor, error)
	Create(ctx context.Context, props *model.Fornecedor) error
	GetByID(ctx context.Context, id int64) (*model.Fornecedor, error)
	Update(ctx context.Context, props *model.Fornecedor) error
	Delete(ctx context.Context, id int64) (*model.Fornecedor, error)
}


func NewHandler(store FornecedorStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /fornecedores", h.getAll)
	mux.HandleFunc("POST /fornecedores", h.create)
	mux.HandleFunc("GET /fornecedores/{id}", h.fetch)
	mux.HandleFunc("PUT /fornecedores/{id}", h.update)
	mux.HandleFunc("DELETE /fornecedores/{id}", h.delete)
}

// @Summary List Fornecedores
// @Tags Fornecedor
// @Produce json
// @Param filter-nome query string false "Filter by nome using operators: like, ilike, eq, ne. Format: operator.value (e.g. like.João)"
// @Param filter-cnpj query string false "Filter by cnpj using operators: eq, ne, like, ilike. Format: operator.value (e.g. eq.123456789)"
// @Param sort query string false "Sort fields: nome, cnpj. Prefix with '-' for desc. Comma separated for multiple fields (e.g. -nome,cnpj)"
// @Param offset query int false "Pagination offset (default 0)"
// @Param limit query int false "Pagination limit (default 10)"
// @Success 200 {array} model.Fornecedor
// @Failure 500 {object} types.ErrorResponse
// @Router /fornecedores [get]
func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	filters, err := NewFornecedorFilter(r.URL.Query())
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
// @Param fornecedor body model.FornecedorCreate true "Fornecedor payload"
// @Success 201 {object} model.Fornecedor
// @Failure 400 {object} types.ErrorResponse
// @Failure 422 {object} types.ErrorResponse
// @Router /fornecedores [post]
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	if r.Body == nil {
		util.ErrorJSON(w, "No body in the request", http.StatusBadRequest)
		return
	}

	var payload model.FornecedorCreate
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := payload.ToFornecedor()
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
// @Failure 400 {object} types.ErrorResponse
// @Failure 404 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /fornecedores/{id} [get]
func (h *Handler) fetch(w http.ResponseWriter, r *http.Request) {
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
// @Param fornecedor body model.FornecedorCreate true "Fornecedor payload"
// @Success 200 {object} model.Fornecedor
// @Failure 400 {object} types.ErrorResponse
// @Failure 422 {object} types.ErrorResponse
// @Router /fornecedores/{id} [put]
func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload model.FornecedorCreate
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := payload.ToFornecedor()
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
// @Failure 400 {object} types.ErrorResponse
// @Failure 422 {object} types.ErrorResponse
// @Router /fornecedores/{id} [delete]
func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
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
