CREATE TABLE IF NOT EXISTS contem_item_oferta (
    quantidade int NOT NULL CHECK (quantidade > 0),
    id_produto int,
    id_oferta int,
    PRIMARY KEY (id_produto, id_oferta),

    FOREIGN KEY (id_produto) REFERENCES ProdutoComercial(id_produto),
    FOREIGN KEY (id_oferta) REFERENCES Oferta(id_oferta) ON DELETE CASCADE
);
