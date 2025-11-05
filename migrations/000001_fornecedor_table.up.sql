CREATE TABLE IF NOT EXISTS Fornecedor (
    CNPJ char(14),
    nome varchar(50) NOT NULL,
    id_fornecedor serial PRIMARY KEY
);
