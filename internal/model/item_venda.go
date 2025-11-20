package model

type ItemVenda struct {
	IDItemVenda   int64   `json:"id_item_venda"`
	IDVenda       int64   `json:"id_venda"`
	IDLote        int64   `json:"id_lote"`
	Quantidade    int64   `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
}

type ItemVendaCreate struct {
	IDVenda       int64   `json:"id_venda"`
	IDLote        int64   `json:"id_lote"`
	Quantidade    int64   `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
}

func (ivc ItemVendaCreate) ToItemVenda() ItemVenda {
	return ItemVenda{
		IDVenda:       ivc.IDVenda,
		IDLote:        ivc.IDLote,
		Quantidade:    ivc.Quantidade,
		ValorUnitario: ivc.ValorUnitario,
	}
}
