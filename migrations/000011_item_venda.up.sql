CREATE TABLE IF NOT EXISTS item_venda (
    id_item_venda SERIAL PRIMARY KEY,
    id_venda int NOT NULL REFERENCES Venda(id_venda) ON DELETE CASCADE,
    id_lote int REFERENCES Lote(id_lote) ON DELETE RESTRICT,
    quantidade int CHECK (quantidade > 0),
    valor_unitario decimal(6, 2) NOT NULL
);
