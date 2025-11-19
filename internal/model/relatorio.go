package model

type SeriePonto struct {
    Date     string  `json:"date"`
    Receita  float64 `json:"receita"`
    Despesa  float64 `json:"despesa"`
    Lucro    float64 `json:"lucro"`
}

type RelatorioFinanceiro struct {
    PeriodStart string       `json:"period_start"`
    PeriodEnd   string       `json:"period_end"`
    Granularity string       `json:"granularity"`
    Totals      struct {
        Receita float64 `json:"receita"`
        Despesa float64 `json:"despesa"`
        Lucro   float64 `json:"lucro"`
    } `json:"totals"`
    Series     []SeriePonto `json:"series"`
    Projection []SeriePonto `json:"projection,omitempty"`
}

type FuncionarioFolhaPagamento struct {
    IdFuncionario   int64   `json:"id_funcionario"`
    Nome            string  `json:"nome"`
    CPF             string  `json:"cpf"`
    Tipo            string  `json:"tipo"`
    Expediente      string  `json:"expediente"`
    SalarioBase     float64 `json:"salario_base"`
    Bonificacao     float64 `json:"bonificacao"`
    SalarioTotal    float64 `json:"salario_total"`
    DataContratacao string  `json:"data_contratacao"`
}

type FolhaPagamentoMensal struct {
    Mes               string                      `json:"mes"`
    Ano               int                         `json:"ano"`
    TotalFuncionarios int                         `json:"total_funcionarios"`
    TotalSalarioBase  float64                     `json:"total_salario_base"`
    TotalBonificacoes float64                     `json:"total_bonificacoes"`
    TotalFolha        float64                     `json:"total_folha"`
    Funcionarios      []FuncionarioFolhaPagamento `json:"funcionarios"`
}

type RelatorioFolhaPagamento struct {
    PeriodStart       string                 `json:"period_start"`
    PeriodEnd         string                 `json:"period_end"`
    TipoFiltro        string                 `json:"tipo_filtro,omitempty"`
    TotalPeriodos     int                    `json:"total_periodos"`
    TotalGeralFolha   float64                `json:"total_geral_folha"`
    FolhasPorMes      []FolhaPagamentoMensal `json:"folhas_por_mes"`
}
