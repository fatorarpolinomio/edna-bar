package relatorio

import (
	"context"
	"net/http"
	"strconv"

	"edna/internal/model"
	"edna/internal/util"
)

type Handler struct {
	store RelatorioStore
}

type RelatorioStore interface {
	GetFinancialReport(ctx context.Context, start, end, granularity string, projectionPeriods int) (model.RelatorioFinanceiro, error)
	GetPayrollReport(ctx context.Context, start, end, tipoFuncionario string) (model.RelatorioFolhaPagamento, error)
}

func NewHandler(store RelatorioStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /relatorios/financeiro", h.getFinancialReport)
	mux.HandleFunc("GET /relatorios/folha-pagamento", h.getPayrollReport)
}

// @Summary Get Financial Report
// @Description Retrieve a financial report within a specified date range.
// @Tags Relatórios
// @Accept json
// @Produce json
// @Param start query string true "Start date (YYYY-MM-DD)"
// @Param end query string true "End date (YYYY-MM-DD)"
// @Param granularity query string false "Time granularity (day|week|month)" default(day)
// @Param projection_days query int false "Number of periods to project"
// @Success 200 {object} model.RelatorioFinanceiro
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /relatorios/financeiro [get]
func (h *Handler) getFinancialReport(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	q := r.URL.Query()
	start := q.Get("start")
	end := q.Get("end")
	granularity := q.Get("granularity")
	projStr := q.Get("projection_days")

	// Basic validation
	if start == "" || end == "" {
		util.ErrorJSON(w, "start and end query parameters are required (YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	projection := 0
	if projStr != "" {
		p, err := strconv.Atoi(projStr)
		if err != nil || p < 0 {
			util.ErrorJSON(w, "projection_days must be a non-negative integer", http.StatusBadRequest)
			return
		}
		projection = p
	}

	// Call store to build the report
	report, err := h.store.GetFinancialReport(ctx, start, end, granularity, projection)
	if err != nil {
		// Return internal server error with the error message
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the report as JSON
	if err := util.WriteJSON(w, http.StatusOK, report); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Get Payroll Report
// @Description Retrieve monthly payroll reports for a specified period. Generates individual payroll for each month within the date range.
// @Tags Relatórios
// @Accept json
// @Produce json
// @Param start query string true "Period start date (YYYY-MM-DD)"
// @Param end query string true "Period end date (YYYY-MM-DD)"
// @Param tipo query string false "Employee type filter (garcom|seguranca|caixa|faxineiro|balconista)"
// @Success 200 {object} model.RelatorioFolhaPagamento
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
// @Router /relatorios/folha-pagamento [get]
func (h *Handler) getPayrollReport(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), util.RequestTimeout)
	defer cancel()

	q := r.URL.Query()
	start := q.Get("start")
	end := q.Get("end")
	tipoFuncionario := q.Get("tipo")

	// Validação básica
	if start == "" || end == "" {
		util.ErrorJSON(w, "start and end query parameters are required (YYYY-MM-DD) - generates monthly payrolls within this period", http.StatusBadRequest)
		return
	}

	// Chamar store para gerar o relatório
	report, err := h.store.GetPayrollReport(ctx, start, end, tipoFuncionario)
	if err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retornar o relatório como JSON
	if err := util.WriteJSON(w, http.StatusOK, report); err != nil {
		util.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
