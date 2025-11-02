package fornecedor

import (
	"context"
	"database/sql"
	"edna/internal/model"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetAll(ctx context.Context) ([]model.Fornecedor, error) {
	query := "SELECT * FROM Fornecedor;"

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var fornecedores []model.Fornecedor
	for rows.Next() {
		var fornecedor model.Fornecedor
		err = rows.Scan(&fornecedor.Id, &fornecedor.Nome, &fornecedor.CNPJ)
		if err != nil {
			return nil, err
		}
		fornecedores = append(fornecedores, fornecedor)
	}

	return fornecedores, nil
}

func (s *Store) Create(ctx context.Context, props *model.Fornecedor) error {
	query := "INSERT INTO Fornecedor (nome, CNPJ) VALUES (?, ?);"

	res, err := s.db.ExecContext(ctx, query, props.Nome, props.CNPJ)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	props.Id = id
	return nil
}
