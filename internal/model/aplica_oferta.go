package model

type AplicaOferta struct {
	IDAplicaOferta int64 `json:"id_aplica_oferta"`
	IDOferta       int64 `json:"id_oferta"`
	IDVenda        int64 `json:"id_venda"`
	IDItemVenda    int64 `json:"id_item_venda"`
}

type AplicaOfertaResponse struct {
	IDOferta    int64 `json:"id_oferta"`
	IDVenda     int64 `json:"id_venda"`
	IDItemVenda int64 `json:"id_item_venda"`
}

func (aor AplicaOfertaResponse) ToAplicaOferta() AplicaOferta {
	return AplicaOferta{
		IDOferta:    aor.IDOferta,
		IDVenda:     aor.IDVenda,
		IDItemVenda: aor.IDItemVenda,
	}
}
