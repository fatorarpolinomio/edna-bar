package model

type Lote struct {
	Id                int64   `json:"id"`
	IdFornecedor      int64   `json:"id_fornecedor"`
	IdProduto         int64   `json:"id_produto"`
	DataFornecimento  string  `json:"data_fornecimento"`
	Validade          string  `json:"validade"`
	PrecoUnitario     float32 `json:"preco_unitario"`
	Estragados        int64   `json:"estragados"`
	QuantidadeInicial int64   `json:"quantidade_inicial"`
}

type LoteCreate struct {
	IdFornecedor      int64   `json:"id_fornecedor"`
	IdProduto         int64   `json:"id_produto"`
	DataFornecimento  string  `json:"data_fornecimento"`
	Validade          string  `json:"validade"`
	PrecoUnitario     float32 `json:"preco_unitario"`
	Estragados        int64   `json:"estragados"`
	QuantidadeInicial int64   `json:"quantidade_inicial"`
}

func (lc LoteCreate) ToLote() Lote {
	return Lote{
		IdFornecedor:      lc.IdFornecedor,
		IdProduto:         lc.IdProduto,
		DataFornecimento:  lc.DataFornecimento,
		Validade:          lc.Validade,
		PrecoUnitario:     lc.PrecoUnitario,
		Estragados:        lc.Estragados,
		QuantidadeInicial: lc.QuantidadeInicial,
	}
}
