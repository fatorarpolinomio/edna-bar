package lote

import (
	"edna/internal/util"
	"net/url"
)

func NewLoteFilter(params url.Values) (util.Filter, error) {
	var filter util.Filter

	if err := filter.GetOffset(params); err != nil {
		return filter, err
	}

	if err := filter.GetLimit(params); err != nil {
		return filter, err
	}

	attrs := []string{"id_fornecedor", "id_produto", "data_fornecimento", "validade", "preco_unitario", "estragados", "quantidade_inicial"}
	if err := filter.GetSorts(params, attrs); err != nil {
		return filter, err
	}

	for _, attr := range attrs {
		if err := filter.GetFilterStr(params, attr); err != nil {
			return filter, err
		}
	}
	return filter, nil
}
