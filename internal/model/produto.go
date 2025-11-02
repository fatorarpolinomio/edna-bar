package model

type Produto struct {
	Id int64 `json:"id"`
	Nome string `json:"nome"`
	Categoria string `json:"categoria"`
	Marca string `json:"marca"`
	QntDisponivel uint `json:"qnt_disponivel"`
	QntTotal uint `json:"qnt_total"`
}

type Estrutural struct {
	Produto
}

type Comercial struct {
	Produto
	PrecoVenda float32 `json:"preco_venda"`
}
