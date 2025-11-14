CREATE TABLE IF NOT EXISTS aplica_oferta (
    id_aplica_oferta SERIAL PRIMARY KEY,
    id_oferta int REFERENCES Oferta(id_oferta) ON DELETE CASCADE,
    id_venda int REFERENCES Venda(id_venda) ON DELETE CASCADE,
    id_item_venda int,

    FOREIGN KEY (id_item_venda) REFERENCES item_venda(id_item_venda)
);
