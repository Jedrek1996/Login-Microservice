package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provide all function to execture db queries and transactions.
// By embedding queries inside store all individual queries will be avaiable to store

// Store - A batch of SQL code that can be used over and over again, dont need to keep calling everyt
type Store struct {
	*Queries
	db *sql.DB
}

// Returns a new store
func newStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

//execTX executes a function within a database transaction

// Create a db transaction, create a new query and call the callback functions of the transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	//New transaction
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx) //New query

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction (tx) err:%v, rollback error:%v", err, rbErr)
		}
		return err
	}
	return tx.Commit() //Commit transaction``
}
