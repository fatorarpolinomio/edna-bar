package produto

import (
	"context"
	"edna/internal/model"
	"edna/internal/types"
	"edna/internal/util"
	"net/http"
)

type Handler struct {
	store ProdutoStore
}

type ProdutoStore interface {
	GetAllComercial(ctx context.Context, filter *model.ComercialFilter) ([]model.Comercial, error)
	GetAllEstrutural(ctx context.Context, filter *model.EstruturalFilter) ([]model.Estrutural, error)
	CreateComercial(ctx context.Context, props *model.Comercial) error
	CreateEstrutural(ctx context.Context, props *model.Estrutural) error
	UpdateComercial(ctx context.Context, id int64, props *model.Comercial) error
	UpdateEstrutural(ctx context.Context, id int64, props *model.Estrutural) error
	GetComercialByID(ctx context.Context, id int64) (*model.Comercial, error)
	GetEstruturalByID(ctx context.Context, id int64) (*model.Estrutural, error)
	Delete(ctx context.Context, id int64) error
}

func NewHandler(store ProdutoStore) Handler {
	return Handler{store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /produtos/comercial", h.getAllComercialHandler)
	mux.HandleFunc("GET /produtos/estrutural", h.getAllEstruturalHandler)
	mux.HandleFunc("POST /produtos/comercial", h.createComercialHandler)
	mux.HandleFunc("POST /produtos/estrutural", h.createEstruturalHandler)
	mux.HandleFunc("PUT /produtos/comercial/{id}", h.updateComercialHandler)
	mux.HandleFunc("PUT /produtos/estrutural/{id}", h.updateEstruturalHandler)
	mux.HandleFunc("GET /produtos/comercial/{id}", h.getComercialHandler)
	mux.HandleFunc("GET /produtos/estrutural/{id}", h.getEstruturalHandler)
	mux.HandleFunc("DELETE /produtos/{id}", h.deleleteProdutoHandler)
}

// @Summary List Comercial Produtos
// @Tags Produtos
// @Produce json
// @Param nome query string false "Filter by nome (partial match)"
// @Param categoria query string false "Filter by categoria"
// @Param marca query string false "Filter by marca"
// @Param offset query int false "Pagination offset (default 0)"
// @Param limit query int false "Pagination limit (default 0)"
// @Param sort query string false "Sort order: asc or desc"
// @Param min-qnt-dsp query int false "Minimum qnt_disponivel"
// @Param max-qnt-dsp query int false "Maximum qnt_disponivel"
// @Param min-qnt-total query int false "Minimum qnt_total"
// @Param max-qnt-total query int false "Maximum qnt_total"
// @Param min-preco-venda query number false "Minimum preco_venda"
// @Param max-preco-venda query number false "Maximum preco_venda"
// @Success 200 {array} model.Comercial
// @Failure 500 {object} map[string]string
// @Router /produtos/comercial [get]
func (h *Handler) getAllComercialHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	filter := model.NewComercialFilter(r.URL)
	produtos, err := h.store.GetAllComercial(ctx, filter)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = util.WriteJSON(w, http.StatusOK, produtos); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

// @Summary List Estrutural Produtos
// @Tags Produtos
// @Produce json
// @Param nome query string false "Filter by nome (partial match)"
// @Param categoria query string false "Filter by categoria"
// @Param marca query string false "Filter by marca"
// @Param offset query int false "Pagination offset (default 0)"
// @Param limit query int false "Pagination limit (default 0)"
// @Param sort query string false "Sort order: asc or desc"
// @Param min-qnt-dsp query int false "Minimum qnt_disponivel"
// @Param max-qnt-dsp query int false "Maximum qnt_disponivel"
// @Param min-qnt-total query int false "Minimum qnt_total"
// @Param max-qnt-total query int false "Maximum qnt_total"
// @Success 200 {array} model.Estrutural
// @Failure 500 {object} map[string]string
// @Router /produtos/estrutural [get]
func (h *Handler) getAllEstruturalHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	filter := model.NewEstruturalFilter(r.URL)
	produtos, err := h.store.GetAllEstrutural(ctx, filter)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = util.WriteJSON(w, http.StatusOK, produtos); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

// @Summary Create Comercial Produto
// @Tags Produtos
// @Accept json
// @Produce json
// @Param produto body model.ProdutoComercialPayload true "Comercial product payload"
// @Success 201 {object} model.Comercial
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /produtos/comercial [post]
func (h *Handler) createComercialHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	payload := model.ProdutoComercialPayload{}
	if err := util.ReadJSON(r, &payload); err != nil {
		util.ErrorJSON(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	produto := model.NewComercial(payload)
	if err := h.store.CreateComercial(ctx, produto); err != nil {
		status := http.StatusInternalServerError
		if err == types.ErrNotFound {
			status = http.StatusNotFound
		}
		util.ErrorJSON(w, err.Error(), status)
		return
	}

	if err := util.WriteJSON(w, http.StatusCreated, produto); err != nil {
		util.ErrorJSON(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Create Estrutural Produto
// @Tags Produtos
// @Accept json
// @Produce json
// @Param produto body model.ProdutoEstruturalPayload true "Estrutural product payload"
// @Success 201 {object} model.Estrutural
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /produtos/estrutural [post]
func (h *Handler) createEstruturalHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	payload := model.ProdutoEstruturalPayload{}
	if err := util.ReadJSON(r, &payload); err != nil {
		util.ErrorJSON(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	produto := model.NewEstrutural(payload)
	if err := h.store.CreateEstrutural(ctx, produto); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := util.WriteJSON(w, http.StatusCreated, produto); err != nil {
		util.ErrorJSON(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Update Comercial Produto
// @Tags Produtos
// @Accept json
// @Produce json
// @Param id path int true "Produto ID"
// @Param produto body model.ProdutoComercialPayload true "Comercial product payload"
// @Success 200 {object} model.Comercial
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /produtos/comercial/{id} [put]
func (h *Handler) updateComercialHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	payload := model.ProdutoComercialPayload{}
	if err := util.ReadJSON(r, &payload); err != nil {
		util.ErrorJSON(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	produto := model.NewComercial(payload)
	if err := h.store.UpdateComercial(ctx, id, produto); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := util.WriteJSON(w, http.StatusOK, produto); err != nil {
		util.ErrorJSON(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Update Estrutural Produto
// @Tags Produtos
// @Accept json
// @Produce json
// @Param id path int true "Produto ID"
// @Param produto body model.ProdutoEstruturalPayload true "Estrutural product payload"
// @Success 200 {object} model.Estrutural
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /produtos/estrutural/{id} [put]
func (h *Handler) updateEstruturalHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	payload := model.ProdutoEstruturalPayload{}
	if err := util.ReadJSON(r, &payload); err != nil {
		util.ErrorJSON(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	produto := model.NewEstrutural(payload)
	if err := h.store.UpdateEstrutural(ctx, id, produto); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := util.WriteJSON(w, http.StatusOK, produto); err != nil {
		util.ErrorJSON(w, "Error encoding response: "+err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Get Comercial Produto by ID
// @Tags Produtos
// @Produce json
// @Param id path int true "Produto ID"
// @Success 200 {object} model.Comercial
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /produtos/comercial/{id} [get]
func (h *Handler) getComercialHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	produto, err := h.store.GetComercialByID(ctx, id)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := util.WriteJSON(w, http.StatusOK, produto); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

// @Summary Get Estrutural Produto by ID
// @Tags Produtos
// @Produce json
// @Param id path int true "Produto ID"
// @Success 200 {object} model.Estrutural
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /produtos/estrutural/{id} [get]
func (h *Handler) getEstruturalHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	produto, err := h.store.GetEstruturalByID(ctx, id)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := util.WriteJSON(w, http.StatusOK, produto); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
	}
}

// @Summary Delete Produto
// @Tags Produtos
// @Param id path int true "Produto ID"
// @Success 204 {string} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /produtos/{id} [delete]
func (h *Handler) deleleteProdutoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	if err := h.store.Delete(ctx, id); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}