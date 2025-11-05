package produto

import (
	"edna/internal/util"
	"net/url"
)


func NewProdutoFilter(params url.Values) (util.Filter, error) {
	var filter util.Filter

	if err := filter.GetOffset(params); err != nil {
		return filter, err
	}

	if err := filter.GetLimit(params); err != nil {
		return filter, err
	}

	attrs := []string{"nome", "categoria", "marca"}
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


func NewComercialFilter(params url.Values) (util.Filter, error) {
	var filter util.Filter

	if err := filter.GetOffset(params); err != nil {
		return filter, err
	}

	if err := filter.GetLimit(params); err != nil {
		return filter, err
	}

	attrs := []string{"nome", "categoria", "marca", "preco_venda"}
	if err := filter.GetSorts(params, attrs); err != nil {
		return filter, err
	}

	for _, attr := range attrs {
		if attr == "preco_venda" {
			if err := filter.GetFilterFloat(params, attr); err != nil {
				return filter, err
			}
		} else {
			if err := filter.GetFilterStr(params, attr); err != nil {
				return filter, err
			}
		}
	}

	return filter, nil
}
