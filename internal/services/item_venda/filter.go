package item_venda

import (
	"edna/internal/util"
	"net/url"
)

func NewItemVendaFilter(params url.Values) (util.Filter, error) {
	var filter util.Filter

	if err := filter.GetOffset(params); err != nil {
		return filter, err
	}

	if err := filter.GetLimit(params); err != nil {
		return filter, err
	}

	attrs := []string{"id_venda", "id_produto", "quantidade", "valor_unitario"}
	if err := filter.GetSorts(params, attrs); err != nil {
		return filter, err
	}

	if err := filter.GetFilterInt(params, "id_venda"); err != nil {
		return filter, err
	}

	if err := filter.GetFilterInt(params, "id_produto"); err != nil {
		return filter, err
	}

	if err := filter.GetFilterInt(params, "quantidade"); err != nil {
		return filter, err
	}

	if err := filter.GetFilterFloat(params, "valor_unitario"); err != nil {
		return filter, err
	}

	return filter, nil
}
