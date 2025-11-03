package model

import (
	"net/url"
	"strconv"
)

type Fornecedor struct {
	Id int64 `json:"id"`
	Nome string `json:"nome"`
	CNPJ string `json:"cnpj"`
}

type FornecedorPayload struct {
	Nome string `json:"nome"`
	CNPJ string `json:"cnpj"`
}

func FromPayload(payload FornecedorPayload) Fornecedor {
	return Fornecedor{
		Nome: payload.Nome,
		CNPJ: payload.CNPJ,
	}
}

type FornecedorFilters struct {
	Nome string
	Offset uint
	Limit uint
	Sort string `enum:"asc,desc"`
}

func NewFornecedorFilter(url *url.URL) FornecedorFilters {
	var filters FornecedorFilters
	filters.Nome = url.Query().Get("nome")
	filters.Sort = url.Query().Get("sort")

	o, err := strconv.ParseUint(url.Query().Get("offset"), 10, 32)
	if err != nil {
		filters.Offset = 0
	} else {
		filters.Offset = uint(o)
	}

	l, err := strconv.ParseUint(url.Query().Get("limit"), 10, 32)
	if err != nil {
		filters.Limit = 10
	} else {
		filters.Limit = uint(l)
	}

	return filters
}
