package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/Xebec19/pathshaala-server/utils"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

func ConnectDB(dsn string) (*DB, error) {
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatalf("cannot load config:", err)
		return nil, err
	}
	const maxOpenDbConn = config.maxOpenDbConn
	const maxIdleDbConn = config.maxIdleDbConn
	const maxDbLifetime = 5 * time.Minute
	d, err := newDatabase(dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
}

func newDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
