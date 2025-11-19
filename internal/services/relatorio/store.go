package relatorio

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"edna/internal/model"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// GetPayrollReport gera um relatório de folha de pagamento mensal para o período especificado
// - start/end são esperados no formato "YYYY-MM-DD" (período do relatório)
// - tipoFuncionario: filtro opcional por tipo de funcionário (garcom, seguranca, caixa, faxineiro, balconista)
// - retorna folhas de pagamento mensais para cada mês dentro do período
func (s *Store) GetPayrollReport(ctx context.Context, start, end, tipoFuncionario string) (model.RelatorioFolhaPagamento, error) {
	var report model.RelatorioFolhaPagamento

	// Validação básica
	if start == "" || end == "" {
		return report, errors.New("start and end são obrigatórios")
	}

	// Parse das datas
	startT, err := time.Parse("2006-01-02", start)
	if err != nil {
		return report, fmt.Errorf("data de início inválida: %w", err)
	}
	endT, err := time.Parse("2006-01-02", end)
	if err != nil {
		return report, fmt.Errorf("data de fim inválida: %w", err)
	}
	if endT.Before(startT) {
		return report, errors.New("data de fim deve ser >= data de início")
	}

	// Gerar folhas mensais
	var folhasMensais []model.FolhaPagamentoMensal
	var totalGeralFolha float64

	// Iterar por cada mês no período
	current := time.Date(startT.Year(), startT.Month(), 1, 0, 0, 0, 0, startT.Location())
	endMonth := time.Date(endT.Year(), endT.Month(), 1, 0, 0, 0, 0, endT.Location())

	for !current.After(endMonth) {
		folhaMensal, err := s.generateMonthlyPayroll(ctx, current, tipoFuncionario)
		if err != nil {
			return report, fmt.Errorf("erro ao gerar folha de %s/%d: %w",
				current.Month().String(), current.Year(), err)
		}

		folhasMensais = append(folhasMensais, folhaMensal)
		totalGeralFolha += folhaMensal.TotalFolha

		// Próximo mês
		current = current.AddDate(0, 1, 0)
	}

	// Montar relatório final
	report.PeriodStart = startT.Format("2006-01-02")
	report.PeriodEnd = endT.Format("2006-01-02")
	report.TipoFiltro = tipoFuncionario
	report.TotalPeriodos = len(folhasMensais)
	report.TotalGeralFolha = totalGeralFolha
	report.FolhasPorMes = folhasMensais

	return report, nil
}

// generateMonthlyPayroll gera a folha de pagamento para um mês específico
func (s *Store) generateMonthlyPayroll(ctx context.Context, month time.Time, tipoFuncionario string) (model.FolhaPagamentoMensal, error) {
	var folha model.FolhaPagamentoMensal

	// Último dia do mês
	lastDay := time.Date(month.Year(), month.Month()+1, 0, 23, 59, 59, 0, month.Location())

	// Query para funcionários ativos no mês (contratados até o último dia do mês)
	query := `
		SELECT
			id_funcionario,
			nome,
			CPF,
			tipo::text,
			expediente::text,
			salario,
			data_contratacao
		FROM Funcionario
		WHERE data_contratacao <= $1::date`

	var args []interface{}
	args = append(args, lastDay.Format("2006-01-02"))

	// Adicionar filtro por tipo se especificado
	if tipoFuncionario != "" {
		query += " AND tipo = $2"
		args = append(args, tipoFuncionario)
	}

	query += " ORDER BY nome"

	// Executar query
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return folha, fmt.Errorf("erro ao consultar funcionários: %w", err)
	}
	defer rows.Close()

	var funcionarios []model.FuncionarioFolhaPagamento
	var totalSalarioBase, totalBonificacoes float64

	for rows.Next() {
		var funcio model.FuncionarioFolhaPagamento
		err := rows.Scan(
			&funcio.IdFuncionario,
			&funcio.Nome,
			&funcio.CPF,
			&funcio.Tipo,
			&funcio.Expediente,
			&funcio.SalarioBase,
			&funcio.DataContratacao,
		)
		if err != nil {
			return folha, fmt.Errorf("erro ao escanear funcionário: %w", err)
		}

		// Calcular bonificação baseada no tipo e salário
		funcio.Bonificacao = s.calculateBonificacao(funcio.Tipo, funcio.SalarioBase)
		funcio.SalarioTotal = funcio.SalarioBase + funcio.Bonificacao

		funcionarios = append(funcionarios, funcio)
		totalSalarioBase += funcio.SalarioBase
		totalBonificacoes += funcio.Bonificacao
	}

	if err := rows.Err(); err != nil {
		return folha, fmt.Errorf("erro durante iteração dos resultados: %w", err)
	}

	// Montar folha mensal
	folha.Mes = month.Month().String()
	folha.Ano = month.Year()
	folha.TotalFuncionarios = len(funcionarios)
	folha.TotalSalarioBase = totalSalarioBase
	folha.TotalBonificacoes = totalBonificacoes
	folha.TotalFolha = totalSalarioBase + totalBonificacoes
	folha.Funcionarios = funcionarios

	return folha, nil
}

// calculateBonificacao calcula bonificação baseada no tipo de funcionário
// Regras simples: garcom e balconista recebem 10%, segurança 15%, outros 5%
func (s *Store) calculateBonificacao(tipo string, salarioBase float64) float64 {
	switch strings.ToLower(tipo) {
	case "garcom", "balconista":
		return salarioBase * 0.10
	case "seguranca":
		return salarioBase * 0.15
	case "caixa", "faxineiro":
		return salarioBase * 0.05
	default:
		return 0.0
	}
}

// GetFinancialReport gera um relatorio financeiro. Com lucro, despesas e ganhos. É possivel definir um intervalo e a granularidade (dia, semana, mes). Bem como, fazer previsões simples com base na média de lucro.
// - start/end are expected in "YYYY-MM-DD" format.
// - granularity: "day", "week", "month"
// - projectionPeriods: number of future periods to project (0 to disable)
func (s *Store) GetFinancialReport(ctx context.Context, start, end, granularity string, projectionPeriods int) (model.RelatorioFinanceiro, error) {
	var report model.RelatorioFinanceiro

	// Basic validation
	if start == "" || end == "" {
		return report, errors.New("start and end are required")
	}
	if granularity == "" {
		granularity = "day"
	}
	if granularity != "day" && granularity != "week" && granularity != "month" {
		return report, errors.New("invalid granularity: must be one of day|week|month")
	}

	// Parse dates
	startT, err := time.Parse("2006-01-02", start)
	if err != nil {
		return report, fmt.Errorf("invalid start date: %w", err)
	}
	endT, err := time.Parse("2006-01-02", end)
	if err != nil {
		return report, fmt.Errorf("invalid end date: %w", err)
	}
	if endT.Before(startT) {
		return report, errors.New("end must be >= start")
	}

	// Fetch aggregations from DB
	receitaMap, err := s.fetchReceita(ctx, start, end, granularity)
	if err != nil {
		return report, fmt.Errorf("fetch receita: %w", err)
	}
	despesaMap, err := s.fetchDespesa(ctx, start, end, granularity)
	if err != nil {
		return report, fmt.Errorf("fetch despesa: %w", err)
	}

	// Build series iterating over periods from start to end
	series := make([]model.SeriePonto, 0)
	totalReceita := 0.0
	totalDespesa := 0.0

	iter := truncateToGranularity(startT, granularity)
	endIter := truncateToGranularity(endT, granularity)

	for !iter.After(endIter) {
		r := receitaMap[iter]
		d := despesaMap[iter]
		l := r - d

		series = append(series, model.SeriePonto{
			Date:    iter.Format(dateFormatForGranularity(granularity)),
			Receita: r,
			Despesa: d,
			Lucro:   l,
		})

		totalReceita += r
		totalDespesa += d

		iter = nextPeriod(iter, granularity)
	}

	// Totals and metadata
	report.PeriodStart = startT.Format("2006-01-02")
	report.PeriodEnd = endT.Format("2006-01-02")
	report.Granularity = granularity
	report.Series = series
	report.Totals.Receita = totalReceita
	report.Totals.Despesa = totalDespesa
	report.Totals.Lucro = totalReceita - totalDespesa

	// Projection (simple): média por periodo
	if projectionPeriods > 0 {
		proj := s.computeProjection(series, projectionPeriods, granularity, endT)
		report.Projection = proj
	}

	return report, nil
}

// fetchReceita retorna a receita de vendas agragado com base em periodo de tempo.
func (s *Store) fetchReceita(ctx context.Context, start, end, granularity string) (map[time.Time]float64, error) {
	agg := make(map[time.Time]float64)

	// Validate granularity -> date_trunc input
	trunc := sqlDateTruncArg(granularity)

	// Build query
	query := fmt.Sprintf(`
	SELECT date_trunc('%s', v.data_hora_venda) AS period,
	       COALESCE(SUM(iv.quantidade * iv.valor_unitario), 0) AS receita
	FROM Venda v
	JOIN item_venda iv ON iv.id_venda = v.id_venda
	WHERE v.data_hora_venda::date BETWEEN $1::date AND $2::date
	GROUP BY period
	ORDER BY period;
`, trunc)

	rows, err := s.db.QueryContext(ctx, query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var period time.Time
		var receita sql.NullFloat64
		if err := rows.Scan(&period, &receita); err != nil {
			return nil, err
		}
		val := 0.0
		if receita.Valid {
			val = receita.Float64
		}
		period = truncateToGranularity(period, granularity)
		agg[period] = val
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return agg, nil
}

// fetchDespesa aggregates expenses based on Lote purchases grouped by the given granularity.
// It computes the purchase cost as preco_unitario * quantidade_inicial grouped by data_fornecimento.
func (s *Store) fetchDespesa(ctx context.Context, start, end, granularity string) (map[time.Time]float64, error) {
	agg := make(map[time.Time]float64)

	trunc := sqlDateTruncArg(granularity)

	query := fmt.Sprintf(`
	SELECT date_trunc('%s', l.data_fornecimento) AS period,
	       COALESCE(SUM(l.preco_unitario * l.quantidade_inicial), 0) AS despesa
	FROM Lote l
	WHERE l.data_fornecimento::date BETWEEN $1::date AND $2::date
	GROUP BY period
	ORDER BY period;
`, trunc)

	rows, err := s.db.QueryContext(ctx, query, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var period time.Time
		var despesa sql.NullFloat64
		if err := rows.Scan(&period, &despesa); err != nil {
			return nil, err
		}
		val := 0.0
		if despesa.Valid {
			val = despesa.Float64
		}
		period = truncateToGranularity(period, granularity)
		agg[period] = val
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return agg, nil
}

// computeProjection creates a simple projection based on average of historical series.
// projectionPeriods is number of future periods to produce (units = granularity).
func (s *Store) computeProjection(series []model.SeriePonto, projectionPeriods int, granularity string, lastDate time.Time) []model.SeriePonto {
	result := make([]model.SeriePonto, 0)
	if len(series) == 0 || projectionPeriods == 0 {
		return result
	}

	// compute averages across available series
	var sumReceita, sumDespesa float64
	for _, p := range series {
		sumReceita += p.Receita
		sumDespesa += p.Despesa
	}
	count := float64(len(series))
	avgReceita := sumReceita / count
	avgDespesa := sumDespesa / count

	// start from the next period after lastDate
	cursor := truncateToGranularity(lastDate, granularity)
	cursor = nextPeriod(cursor, granularity)

	for range projectionPeriods {
		lucro := avgReceita - avgDespesa
		result = append(result, model.SeriePonto{
			Date:    cursor.Format(dateFormatForGranularity(granularity)),
			Receita: avgReceita,
			Despesa: avgDespesa,
			Lucro:   lucro,
		})
		cursor = nextPeriod(cursor, granularity)
	}

	return result
}

// Helpers

// sqlDateTruncArg returns the argument for postgres date_trunc
func sqlDateTruncArg(granularity string) string {
	switch granularity {
	case "day":
		return "day"
	case "week":
		return "week"
	case "month":
		return "month"
	default:
		return "day"
	}
}

// truncateToGranularity truncates a time.Time to the requested granularity.
func truncateToGranularity(t time.Time, granularity string) time.Time {
	switch granularity {
	case "day":
		y, m, d := t.Date()
		return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	case "week":
		// ISO week: find Monday of the week
		weekday := int(t.Weekday())
		// convert Sunday(0) to 7
		if weekday == 0 {
			weekday = 7
		}
		// subtract days to get Monday
		monday := t.AddDate(0, 0, -(weekday - 1))
		y, m, d := monday.Date()
		return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	case "month":
		y, m, _ := t.Date()
		return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
	default:
		return truncateToGranularity(t, "day")
	}
}

// nextPeriod advances the time by one granularity unit.
func nextPeriod(t time.Time, granularity string) time.Time {
	switch granularity {
	case "day":
		return t.AddDate(0, 0, 1)
	case "week":
		return t.AddDate(0, 0, 7)
	case "month":
		return t.AddDate(0, 1, 0)
	default:
		return t.AddDate(0, 0, 1)
	}
}

// dateFormatForGranularity returns a human-friendly date format string for the period label.
func dateFormatForGranularity(granularity string) string {
	switch granularity {
	case "day":
		return "2006-01-02"
	case "week":
		// represent week by starting date (Monday)
		return "2006-01-02"
	case "month":
		return "2006-01"
	default:
		return "2006-01-02"
	}
}
