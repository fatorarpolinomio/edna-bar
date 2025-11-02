/* rollback Lógico_1: */

ALTER TABLE Fornecedor DROP CONSTRAINT IF EXISTS FK_Fornecedor_2;
ALTER TABLE Cliente DROP CONSTRAINT IF EXISTS FK_Cliente_2;
ALTER TABLE Funcionario DROP CONSTRAINT IF EXISTS FK_Funcionario_2;
ALTER TABLE contem_item_oferta DROP CONSTRAINT IF EXISTS FK_contem_item_oferta_2;
ALTER TABLE contem_item_oferta DROP CONSTRAINT IF EXISTS FK_contem_item_oferta_3;
ALTER TABLE contem_item_venda DROP CONSTRAINT IF EXISTS FK_contem_item_venda_2;
ALTER TABLE contem_item_venda DROP CONSTRAINT IF EXISTS FK_contem_item_venda_3;
ALTER TABLE contem_item_lote DROP CONSTRAINT IF EXISTS FK_contem_item_lote_2;
ALTER TABLE contem_item_lote DROP CONSTRAINT IF EXISTS FK_contem_item_lote_3;
ALTER TABLE aplica DROP CONSTRAINT IF EXISTS FK_aplica_1;
ALTER TABLE aplica DROP CONSTRAINT IF EXISTS FK_aplica_2;

DROP TABLE IF EXISTS aplica;
DROP TABLE IF EXISTS contem_item_lote;
DROP TABLE IF EXISTS contem_item_venda;
DROP TABLE IF EXISTS contem_item_oferta;
DROP TABLE IF EXISTS Funcionario;
DROP TABLE IF EXISTS Cliente;
DROP TABLE IF EXISTS Fornecedor;
DROP TABLE IF EXISTS Venda;
DROP TABLE IF EXISTS Lote;
DROP TABLE IF EXISTS Oferta;
DROP TABLE IF EXISTS Produto;
