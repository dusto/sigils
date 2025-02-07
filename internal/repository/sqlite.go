package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

var DefaultConnectionParams = func() *url.Values {
	params := &url.Values{}

	// Ensure transactions lock immediately
	params.Add("_txlock", "immediate")

	// Enable support for foreign_keys
	params.Add("_foreign_keys", "true")

	// Enables WAL to allow for concurrent readers and writers
	params.Add("_journal_mode", "WAL")

	// Set timeout to acquire a lock to 5 seconds
	params.Add("_busy_timeout", "5000")

	// Allow database to sync but less often
	params.Add("_synchronous", "NORMAL")

	// Increase cache size: based on 4096byte page size this sets the cache size to about 4Gb
	params.Add("_cache_size", "1000000")

	return params
}

// MultiSqliteDB is a wrapper for sqlite that enables doing read and writes at the same time
type MultiSqliteDB struct {
	writeDB *sql.DB
	readDB  *sql.DB
}

func (db *MultiSqliteDB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.writeDB.ExecContext(ctx, query, args...)
}

func (db *MultiSqliteDB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.readDB.PrepareContext(ctx, query)
}

func (db *MultiSqliteDB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return db.readDB.QueryContext(ctx, query, args...)
}

func (db *MultiSqliteDB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return db.readDB.QueryRowContext(ctx, query, args...)

}

func (db *MultiSqliteDB) BeginWriteTx(ctx context.Context) (*sql.Tx, error) {
	return db.writeDB.BeginTx(ctx, nil)

}

func (db *MultiSqliteDB) SetupMultiSqliteDB(path string, connectionParams *url.Values) error {
	connectUri := fmt.Sprintf("file:%s?%s", path, connectionParams.Encode())
	var err error
	db.writeDB, err = sql.Open("sqlite3", connectUri)
	if err != nil {
		return err
	}
	// Only allow one connection for writing since only 1 write can happen at a time with sqlite
	db.writeDB.SetMaxOpenConns(1)

	db.readDB, err = sql.Open("sqlite3", connectUri)
	if err != nil {
		return err
	}

	// We can have multiple read connections
	db.readDB.SetMaxOpenConns(max(4, runtime.NumCPU()))

	return nil
}
