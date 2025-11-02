package model

type Fornecedor struct {
	Id int64 `json:"id"`
	Nome string `json:"nome"`
	CNPJ string `json:"cpnj"`
}
