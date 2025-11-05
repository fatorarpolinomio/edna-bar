package model


type Fornecedor struct {
	Id int64 `json:"id"`
	Nome string `json:"nome"`
	CNPJ string `json:"cnpj"`
}

type FornecedorCreate struct {
	Nome string `json:"nome"`
	CNPJ string `json:"cnpj"`
}

func (fc FornecedorCreate) ToFornecedor() Fornecedor {
	return Fornecedor{
		Nome: fc.Nome,
		CNPJ: fc.CNPJ,
	}
}
