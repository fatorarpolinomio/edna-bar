package produto

import (
	"context"
	"database/sql"
	"edna/internal/model"
	"edna/internal/types"
	"edna/internal/util"
	"log"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetAll(ctx context.Context, filter *util.Filter) ([]model.UnionProduto, error) {
	query := "SELECT p.id_produto, p.nome, p.categoria, p.marca, c.preco_venda FROM Produto p LEFT JOIN ProdutoComercial AS c using (id_produto)"
	rows, err := util.QueryRowsWithFilter(s.db, ctx, query, filter, "p")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produtos := make([]model.UnionProduto, 0)
	for rows.Next() {
		c := model.UnionProduto{}
		err = rows.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca, &c.PrecoVenda)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, types.ErrNotFound
			}
			log.Printf("Error scanning row: %v", err)
			return nil, types.ErrInternalServer
		}
		produtos = append(produtos, c)
	}

	return produtos, nil
}

func (s *Store) GetAllComercial(ctx context.Context, filter *util.Filter) ([]model.Comercial, error) {
	query := "SELECT id_produto, nome, categoria, marca, preco_venda FROM ProdutoComercial c"
	rows, err := util.QueryRowsWithFilter(s.db, ctx, query, filter, "c")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produtos := make([]model.Comercial, 0)
	for rows.Next() {
		c := model.Comercial{}
		err = rows.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca, &c.PrecoVenda)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, types.ErrNotFound
			}
			log.Printf("Error scanning row: %v", err)
			return nil, types.ErrInternalServer
		}
		produtos = append(produtos, c)
	}

	return produtos, nil
}


func (s *Store) GetAllEstrutural(ctx context.Context, filter *util.Filter) ([]model.Produto, error) {
	query := "SELECT id_produto, nome, categoria, marca FROM ONLY Produto p"
	rows, err := util.QueryRowsWithFilter(s.db, ctx, query, filter, "p")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	produtos := make([]model.Produto, 0)
	for rows.Next() {
		c := model.Produto{}
		err = rows.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, types.ErrNotFound
			}
			log.Printf("Error scanning row: %v", err)
			return nil, types.ErrInternalServer
		}
		produtos = append(produtos, c)
	}

	return produtos, nil
}


func (s *Store) CreateComercial(ctx context.Context, props *model.Comercial) error {
	// Executa uma transação (desfaz insercao em caso de erro) pois precisamos fazer 2 insercoes
	query := "INSERT INTO ProdutoComercial (nome, categoria, marca, preco_venda) VALUES ($1, $2, $3, $4) RETURNING id_produto;"
	row := s.db.QueryRowContext(ctx, query, props.Nome, props.Categoria, props.Marca, props.PrecoVenda)
	// Atualiza o id do produto
	err := row.Scan(&props.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Create(ctx context.Context, props *model.Produto) error {
	query := "INSERT INTO Produto (nome, categoria, marca) VALUES ($1, $2, $3) RETURNING id_produto;"

	row := s.db.QueryRowContext(ctx, query, props.Nome, props.Categoria, props.Marca)
	err := row.Scan(&props.Id)
	if err != nil {
		return err
	}
	return nil
}
func (s *Store) UpdateComercial(ctx context.Context, props *model.Comercial) error {
	query := "UPDATE ProdutoComercial SET nome = $1, categoria = $2, marca = $3, preco_venda = $4 WHERE id_produto = $5;"
	res, err := s.db.ExecContext(ctx, query, props.Nome, props.Categoria, props.Marca, props.PrecoVenda, props.Id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return types.ErrNotFound
	}
	return nil
}

func (s *Store) Update(ctx context.Context, props *model.Produto) error {
	query := "UPDATE Produto SET nome = $1, categoria = $2, marca = $3, quantidade_disponivel = $4, quantidade_total = $5 WHERE id_produto = $6;"

	res, err := s.db.ExecContext(ctx, query, props.Nome, props.Categoria, props.Marca, props.Id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return types.ErrNotFound
	}
	return nil
}

func (s *Store) GetComercialByID(ctx context.Context, id int64) (*model.Comercial, error) {
	query := "SELECT id_produto, nome, categoria, marca, preco_venda FROM ProdutoComercial WHERE id_produto = $1"
	row := s.db.QueryRowContext(ctx, query, id)
	c := model.Comercial{}
	err := row.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca, &c.PrecoVenda)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, types.ErrNotFound
		case sql.ErrConnDone:
			log.Println("[WARN] Connection closed.")
			return nil, types.ErrInternalServer
		default:
			log.Println("[ERROR] Unexpected error:", err)
			return nil, err
		}
	}
	return &c, nil
}

func (s *Store) GetByID(ctx context.Context, id int64) (*model.Produto, error) {
	query := "SELECT id_produto, nome, categoria, marca FROM Produto WHERE id_produto = $1"
	row := s.db.QueryRowContext(ctx, query, id)
	c := model.Produto{}
	err := row.Scan(&c.Id, &c.Nome, &c.Categoria, &c.Marca)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (s *Store) Delete(ctx context.Context, id int64) error {
	// Derivadas do produto serão apagadas automaticamente por conta da herança
	query := "DELETE FROM Produto WHERE id_produto = $1"
	_, err := s.db.ExecContext(ctx, query, id)
	return err
}
