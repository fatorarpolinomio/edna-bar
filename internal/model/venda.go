package model

import (
	"time"
)

type Venda struct {
	Id                int64     `json:"id"`
	IdCliente         int64     `json:"id_cliente"`
	IdFuncionario     int64     `json:"id_funcionario"`
	DataHoraVenda     time.Time `json:"data_hora_renda"`
	DataHoraPagamento *time.Time `json:"data_hora_pagamento"`
	TipoPagamento     string    `json:"tipo_pagamento"`
}

type VendaCreate struct {
	IdCliente         int64     `json:"id_cliente"`
	IdFuncionario     int64     `json:"id_funcionario"`
	DataHoraVenda     time.Time `json:"data_hora_renda"`
	DataHoraPagamento *time.Time `json:"data_hora_pagamento"`
	TipoPagamento     string    `json:"tipo_pagamento"`
}

func (vc *VendaCreate) ToVenda() Venda {
	return Venda{
		IdCliente:         vc.IdCliente,
		IdFuncionario:     vc.IdFuncionario,
		DataHoraVenda:     vc.DataHoraVenda,
		DataHoraPagamento: vc.DataHoraPagamento,
		TipoPagamento:     vc.TipoPagamento,
	}
}
