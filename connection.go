package ldb

import (
	"context"
	"database/sql"
	"log"
)

type (
	Connection struct {
		Configuration *Configuration
		Database      *sql.DB
	}

	Parse[T any] func(*sql.Rows) (T, error)
)

// Connect to a database with given configurations
func Connect(config *Configuration) *Connection {
	db, err := sql.Open(config.Driver, config.ConnectionString())
	if err != nil {
		log.Fatalf("Error loading database from given configuration due to %s", err)
	}

	return &Connection{
		Configuration: config,
		Database:      db,
	}
}

// Q performs a query to the database
func Q[T any](conn *Connection, query string, parameters []any, parse Parse[T]) ([]T, error) {
	rows, err := conn.Database.QueryContext(context.TODO(), query, parameters...)
	if err != nil {
		return nil, err
	}
	// TODO: Add fine-tuned logging for error cases and defers
	defer rows.Close()

	return parseRows(rows, parse)
}

// Tx starts a transaction in the database where all transaction queries happen in body
func Tx[T any](conn *Connection, body func(*sql.Tx) (T, error)) (T, error) {
	tx, err := conn.Database.BeginTx(context.TODO(), nil)
	if err != nil {
		return *new(T), err
	}
	// TODO: Handle if error rolling back
	defer tx.Rollback()

	result, err := body(tx)
	if err != nil {
		// Any of the transaction queries fail
		return *new(T), err
	}

	err = tx.Commit()
	if err != nil {
		// Failed to commit
		return *new(T), err
	}

	return result, nil
}

// TxQ performs a query in a transaction
func TxQ[T any](tx *sql.Tx, query string, parameters []any, parse Parse[T]) ([]T, error) {
	rows, err := tx.QueryContext(context.TODO(), query, parameters...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseRows(rows, parse)
}

// P generates the parameters for a parameterised query
func P(args ...any) []any {
	return args
}

func parseRows[T any](rows *sql.Rows, fn Parse[T]) ([]T, error) {
	var items []T

	for rows.Next() {
		item, err := fn(rows)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
