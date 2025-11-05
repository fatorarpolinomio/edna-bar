-- Script SQL de referência
-- Veja em migrations/ para ter uma visão completa

CREATE TABLE IF NOT EXISTS Fornecedor (
    CNPJ char(14),
    nome varchar(50) NOT NULL,
    id_fornecedor serial PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS Cliente (
    id_cliente serial PRIMARY KEY,
    CPF char(11),
    data_nascimento date,
    nome varchar(50) NOT NULL
);

DROP TYPE IF EXISTS tipo_de_funcionario;
CREATE TYPE tipo_de_funcionario AS ENUM ('garcom', 'seguranca',
'caixa', 'faxineiro', 'balconista');

DROP TYPE IF EXISTS tipo_de_expediente;
CREATE TYPE tipo_de_expediente AS ENUM ('manha', 'tarde', 'noite', 'madrugada');

CREATE TABLE IF NOT EXISTS Funcionario (
    id_funcionario serial PRIMARY KEY,
    nome varchar(50) NOT NULL,
    CPF char(11) NOT NULL,
    tipo tipo_de_funcionario NOT NULL,
    expediente tipo_de_expediente NOT NULL,
    data_contratacao date NOT NULL,
    salario decimal(8, 2) NOT NULL
);

DROP TYPE IF EXISTS tipo_de_produto;
CREATE TYPE tipo_de_produto AS ENUM ('estrutural', 'comercial');

CREATE TABLE IF NOT EXISTS Produto (
    id_produto serial PRIMARY KEY,
    nome varchar(50) NOT NULL,
    categoria text,
    marca text,
);

CREATE TABLE IF NOT EXISTS ProdutoComercial (
    preco_venda decimal(6, 2) CHECK (VALUE > 0) NOT NULL,
) INHERITS (Produto);

CREATE TABLE IF NOT EXISTS Lote (
    id_lote serial PRIMARY KEY,
    id_fornecedor int NOT NULL,
    id_produto int NOT NULL,
    data_fornecimento date NOT NULL,
    validate date,
    preco_unitario decimal(6, 2) NOT NULL CHECK (VALUE > 0),
    estragados int CHECK (VALUE >= 0),
    qnt_comprada int CHECK (VALUE > 0),


    FOREIGN KEY (id_fornecedor) REFERENCES Fornecedor(id_fornecedor),
    FOREIGN KEY (id_produto) REFERENCES Produto(id_produto)
);

CREATE TABLE IF NOT EXISTS Oferta (
    id_oferta serial PRIMARY KEY,
    nome varchar(50) NOT NULL,
    data_criacao date NOT NULL DEFAULT CURRENT_DATE,
    data_inicio date,
    data_fim date,
    valor_fixo decimal(6, 2),
    percentual_desconto int
);

CREATE TABLE IF NOT EXISTS item_oferta (
    id_item_oferta serial PRIMARY KEY,
    id_oferta int REFERENCES Oferta(id_oferta) ON DELETE CASCADE,
    id_produto int REFERENCES ProdutoComercial(id_produto) ON DELETE RESTRICT
);

DROP TYPE IF EXISTS tipo_de_pagamento;
CREATE TYPE tipo_de_pagamento AS ENUM ('credito', 'debito', 'pix', 'dinheiro', 'VA/VR');

CREATE TABLE IF NOT EXISTS Venda (
    id_venda serial PRIMARY KEY,
    data_hora_venda timestamp NOT NULL DEFAULT now(),
    data_hora_pagamento timestamp,
    tipo_pagamento tipo_de_pagamento,

    id_cliente int NOT NULL,
    id_funcionario int NOT NULL,

    FOREIGN KEY (id_cliente) REFERENCES Cliente(id_cliente) ON DELETE RESTRICT,
    FOREIGN KEY (id_funcionario) REFERENCES Funcionario(id_funcionario) ON DELETE SET NULL
);


CREATE TABLE IF NOT EXISTS item_venda (
    id_item_venda SERIAL PRIMARY KEY,
    id_venda int NOT NULL REFERENCES Venda(id_venda) ON DELETE CASCADE,
    data_hora_venda timestamp NOT NULL DEFAULT now(),
    id_lote int REFERENCES Lote(id_lote) ON DELETE RESTRICT,
    quantidade int CHECK (VALUE > 0),
    valor_unitario decimal(6, 2) NOT NULL,
);

CREATE TABLE IF NOT EXISTS aplica_oferta (
    id_aplica_oferta SERIAL PRIMARY KEY,
    id_venda int REFERENCES Venda(id_venda) ON DELETE CASCADE,
    id_oferta int REFERENCES Oferta(id_oferta) ON DELETE CASCADE,
    id_item_venda int,

    FOREIGN KEY (id_item_venda) REFERENCES item_venda(id_item_venda)
);
