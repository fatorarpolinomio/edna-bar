package item_venda

import (
	"context"
	"edna/internal/model"
	"edna/internal/types"
	"edna/internal/util"
	"encoding/json"
	"net/http"
)

type Handler struct {
	store ItemVendaStore
}

type ItemVendaStore interface {
	GetAll(ctx context.Context, filter util.Filter) ([]model.ItemVenda, error)
	Create(ctx context.Context, props *model.ItemVenda) error
	GetByID(ctx context.Context, id int64) (*model.ItemVenda, error)
	Update(ctx context.Context, props *model.ItemVenda) error
	Delete(ctx context.Context, id int64) (*model.ItemVenda, error)
}

func NewHandler(store ItemVendaStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /item_venda", h.getAll)
	mux.HandleFunc("POST /item_venda", h.create)
	mux.HandleFunc("GET /item_venda/{id}", h.fetch)
	mux.HandleFunc("PUT /item_venda/{id}", h.update)
	mux.HandleFunc("DELETE /item_venda/{id}", h.delete)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	filters, err := NewItemVendaFilter(r.URL.Query())
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	itensVenda, err := h.store.GetAll(ctx, filters)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = util.WriteJSON(w, http.StatusOK, itensVenda)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	if r.Body == nil {
		util.ErrorJSON(w, "No body in the request", http.StatusBadRequest)
		return
	}

	var payload model.ItemVendaCreate
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := payload.ToItemVenda()
	err = h.store.Create(ctx, &model)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.WriteJSON(w, http.StatusCreated, model)
}

func (h *Handler) fetch(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemVenda, err := h.store.GetByID(ctx, id)
	if err != nil {
		if err == types.ErrNotFound {
			util.ErrorJSON(w, "ItemVenda not found.", http.StatusNotFound)
			return
		}
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = util.WriteJSON(w, http.StatusOK, itemVenda); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	id, err := util.GetIDParam(r)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload model.ItemVendaCreate
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	model := payload.ToItemVenda()
	model.IDItemVenda = id
	err = h.store.Update(ctx, &model)
	if err != nil {
		if err == types.ErrNotFound {
			util.ErrorJSON(w, "ItemVenda not found.", http.StatusNotFound)
			return
		}
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.WriteJSON(w, http.StatusOK, model)
}

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
		if err == types.ErrNotFound {
			util.ErrorJSON(w, "ItemVenda not found.", http.StatusNotFound)
			return
		}
		util.ErrorJSON(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	util.WriteJSON(w, http.StatusOK, model)
}
