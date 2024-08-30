package storage

import (
	"database/sql"
	"log"
)

func BeginTransaction(db *sql.DB) (*sql.Tx, error) {
    tx, err := db.Begin()
    if err != nil {
        return nil, err
    }
    return tx, nil
}

func CommitTransaction(tx *sql.Tx) error {
    if err := tx.Commit(); err != nil {
        return err
    }
    return nil
}

func RollbackTransaction(tx *sql.Tx) {
    if err := tx.Rollback(); err != nil {
        log.Printf("Transaction rollback failed: %v", err)
    }
}
