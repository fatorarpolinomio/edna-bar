package model

type ItemOferta struct {
	Quantidade int64 `json:"quantidade"`
	IDProduto  int64 `json:"id_produto"`
	IDOferta   int64 `json:"id_oferta"`
}

type ItemOfertaCreate struct {
	Quantidade int64 `json:"quantidade"`
	IDProduto  int64 `json:"id_produto"`
	IDOferta   int64 `json:"id_oferta"`
}

func (ioc ItemOfertaCreate) ToItemOferta() ItemOferta {
	return ItemOferta{
		Quantidade: ioc.Quantidade,
		IDProduto:  ioc.IDProduto,
		IDOferta:   ioc.IDOferta,
	}
}
