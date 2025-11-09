package model

type Cliente struct {
	Id         int64  `json:"id"`
	CPF        string `json:"cpf"`
	Nascimento string `json:"data_nascimento"`
	Nome       string `json:"nome"`
}

type ClienteCreate struct {
	Nome       string `json:"nome"`
	Nascimento string `json:"data_nascimento"`
	CPF        string `json:"cnpj"`
}

func (cc ClienteCreate) ToCliente() Cliente {
	return Cliente{
		Nome:       cc.Nome,
		Nascimento: cc.Nascimento,
		CPF:        cc.CPF,
	}
}
