CREATE TABLE IF NOT EXISTS Lote (
    id_lote serial PRIMARY KEY,
    id_fornecedor int NOT NULL,
    id_produto int NOT NULL,

    data_fornecimento date NOT NULL,
    validade date,
    preco_unitario decimal(6, 2) NOT NULL CHECK (preco_unitario > 0),
    estragados int CHECK (estragados >= 0) DEFAULT 0,
    quantidade_inicial int CHECK (quantidade_inicial > 0),

    FOREIGN KEY (id_fornecedor) REFERENCES Fornecedor(id_fornecedor),
    FOREIGN KEY (id_produto) REFERENCES Produto(id_produto)
);
