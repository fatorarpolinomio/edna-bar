package lote

import (
	"context"
	"database/sql"
	"edna/internal/model"
	"edna/internal/types"
	"edna/internal/util"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetAll(ctx context.Context, filter util.Filter) ([]model.Cliente, error) {
	query := "SELECT id_cliente, nome, CPF, data_nascimento  FROM Cliente AS c"

	rows, err := util.QueryRowsWithFilter(s.db, ctx, query, &filter, "f")
	if err != nil {
		return nil, err
	}

	clientes := make([]model.Cliente, 0)
	for rows.Next() {
		var cliente model.Cliente
		err = rows.Scan(&cliente.Id, &cliente.Nome, &cliente.CPF)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

func (s *Store) Create(ctx context.Context, props *model.Cliente) error {
	query := "INSERT INTO Cliente (nome, CPF, data_nascimento) VALUES ($1, $2, $3) RETURNING id_cliente;"

	res := s.db.QueryRowContext(ctx, query, props.Nome, props.CPF, props.Nascimento)
	return res.Scan(&props.Id)
}

func (s *Store) GetByID(ctx context.Context, id int64) (*model.Cliente, error) {
	query := "SELECT id_cliente, nome, CPF, data_nascimento FROM Cliente WHERE id_cliente = $1;"

	row := s.db.QueryRowContext(ctx, query, id)

	var cliente model.Cliente
	err := row.Scan(&cliente.Id, &cliente.Nome, &cliente.CPF)
	if err != nil {
		return nil, err
	}

	return &cliente, nil
}

func (s *Store) Update(ctx context.Context, props *model.Cliente) error {
	query := "UPDATE Cliente SET nome = $1, CPF = $2, data_nascimento = $3 WHERE id_cliente = $3;"

	res, err := s.db.ExecContext(ctx, query, props.Nome, props.CPF, props.Id, props.Nascimento)
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

func (s *Store) Delete(ctx context.Context, id int64) (*model.Cliente, error) {
	query := "DELETE FROM Cliente WHERE id_cliente = $1 RETURNING id_cliente, nome, CPF, data_nascimento;"

	var model model.Cliente
	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&model.Id, &model.Nome, &model.CPF)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
