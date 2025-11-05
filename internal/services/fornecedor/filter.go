package fornecedor

import (
	"edna/internal/util"
	"net/url"
)


func NewFornecedorFilter(params url.Values) (util.Filter, error) {
	var filter util.Filter

	if err := filter.GetOffset(params); err != nil {
		return filter, err
	}

	if err := filter.GetLimit(params); err != nil {
		return filter, err
	}

	attrs := []string{"nome", "cnpj"}
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
