/* Lógico_1: */

CREATE TABLE Produto (
    id_produto int serial PRIMARY KEY,
    nome varchar(50) NOT NULL,
    categoria text,
    marca text,
    quantidade_disponivel int NOT NULL,
    quantidade_total int NOT NULL,
    preco_venda decimal(6, 2),
    Produto_TIPO INT
);

CREATE TABLE Fornecedor (
    CNPJ char(11),
    nome varchar(50) NOT NULL,
    id_fornecedor int PRIMARY KEY,
    fk_Lote_id_lote int
);

CREATE TABLE Lote (
    id_lote int serial PRIMARY KEY,
    data_fornecimento date NOT NULL
);

CREATE TABLE Venda (
    id_venda int serial PRIMARY KEY,
    data_hora_venda datetime,
    data_hora_pagamento datetime,
    tipo_pagamento int
);

CREATE TABLE Cliente (
    id_cliente int serial PRIMARY KEY,
    CPF char(11),
    data_nascimento date,
    nome varchar(50) NOT NULL,
    fk_Venda_id_venda int
);

CREATE TABLE Oferta (
    id_oferta int serial PRIMARY KEY,
    nome varchar(50) NOT NULL,
    data_criacao date,
    data_inicio date,
    data_fim date,
    valor_fixo decimal(6, 2),
    percentual_desconto int
);

CREATE TABLE Funcionario (
    id_funcionario int serial PRIMARY KEY,
    nome varchar(50) NOT NuLL,
    CPF char(11) NOT NULL,
    tipo ENUM(garcom, seguranca, caixa, faxineiro, balconista) NOT NULL,
    expediente int NOT NULL,
    data_contratacao date NOT NULL,
    salario decimal(8, 2) NOT NULL,
    fk_Venda_id_venda int
);

CREATE TABLE contem_item_oferta (
    quantidade int NOT NULL,
    id_produto int,
    id_oferta int,
    fk_Oferta_id_oferta int,
    fk_Produto_id_produto int,
    PRIMARY KEY (id_produto, id_oferta)
);

CREATE TABLE contem_item_venda (
    id_oferta int,
    valor_unitario decimal(6, 2),
    quantidade int,
    id_produto int,
    id_venda int,
    fk_Venda_id_venda int,
    fk_Produto_id_produto int,
    PRIMARY KEY (id_produto, id_venda)
);

CREATE TABLE contem_item_lote (
    quantidade_utilizavel int,
    id_lote int,
    quantidade_manutencao int,
    id_produto int,
    quantidade_quebrada int,
    quantidade_comprada int NOT NULL,
    preco_unitario decimal(6, 2) NOT NULL,
    validade date,
    quantidade_disponivel int,
    quantidade_armazenada int,
    quantidade_estragada int,
    fk_Lote_id_lote int,
    fk_Produto_id_produto int,
    PRIMARY KEY (id_lote, id_produto)
);

CREATE TABLE aplica (
    fk_Venda_id_venda int,
    fk_Oferta_id_oferta int
);

ALTER TABLE Fornecedor ADD CONSTRAINT FK_Fornecedor_2
    FOREIGN KEY (fk_Lote_id_lote)
    REFERENCES Lote (id_lote)
    ON DELETE RESTRICT;

ALTER TABLE Cliente ADD CONSTRAINT FK_Cliente_2
    FOREIGN KEY (fk_Venda_id_venda)
    REFERENCES Venda (id_venda)
    ON DELETE RESTRICT;

ALTER TABLE Funcionario ADD CONSTRAINT FK_Funcionario_2
    FOREIGN KEY (fk_Venda_id_venda)
    REFERENCES Venda (id_venda)
    ON DELETE CASCADE;

ALTER TABLE contem_item_oferta ADD CONSTRAINT FK_contem_item_oferta_2
    FOREIGN KEY (fk_Oferta_id_oferta)
    REFERENCES Oferta (id_oferta);

ALTER TABLE contem_item_oferta ADD CONSTRAINT FK_contem_item_oferta_3
    FOREIGN KEY (fk_Produto_id_produto)
    REFERENCES Produto (id_produto);

ALTER TABLE contem_item_venda ADD CONSTRAINT FK_contem_item_venda_2
    FOREIGN KEY (fk_Venda_id_venda)
    REFERENCES Venda (id_venda);

ALTER TABLE contem_item_venda ADD CONSTRAINT FK_contem_item_venda_3
    FOREIGN KEY (fk_Produto_id_produto)
    REFERENCES Produto (id_produto);

ALTER TABLE contem_item_lote ADD CONSTRAINT FK_contem_item_lote_2
    FOREIGN KEY (fk_Lote_id_lote)
    REFERENCES Lote (id_lote);

ALTER TABLE contem_item_lote ADD CONSTRAINT FK_contem_item_lote_3
    FOREIGN KEY (fk_Produto_id_produto)
    REFERENCES Produto (id_produto);

ALTER TABLE aplica ADD CONSTRAINT FK_aplica_1
    FOREIGN KEY (fk_Venda_id_venda)
    REFERENCES Venda (id_venda)
    ON DELETE SET NULL;

ALTER TABLE aplica ADD CONSTRAINT FK_aplica_2
    FOREIGN KEY (fk_Oferta_id_oferta)
    REFERENCES Oferta (id_oferta)
    ON DELETE SET NULL;
