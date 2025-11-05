package model

type Produto struct {
	Id int64 `json:"id"`
	Nome string `json:"nome"`
	Categoria string `json:"categoria"`
	Marca string `json:"marca"`
}

type Comercial struct {
	Produto
	PrecoVenda float32 `json:"preco_venda"`
}

// Uniao entre produto estrutural e comercial
type UnionProduto struct {
	Produto
	PrecoVenda *float32 `json:"preco_venda"`
}

type ProdutoCreate struct {
	Nome string `json:"nome"`
	Categoria string `json:"categoria"`
	Marca string `json:"marca"`
}

type ComercialCreate struct {
	ProdutoCreate
	PrecoVenda float32 `json:"preco_venda"`
}


func (pc ProdutoCreate) ToProduto() Produto {
	return Produto{
		Nome: pc.Nome,
		Categoria: pc.Categoria,
		Marca: pc.Marca,
	}
}

func (cc ComercialCreate) ToComercial() Comercial {
	return Comercial{
		Produto: cc.ProdutoCreate.ToProduto(),
		PrecoVenda: cc.PrecoVenda,
	}
}
