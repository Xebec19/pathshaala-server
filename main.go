package main

import (
	"database/sql"
	"log"

	db "github.com/Xebec19/pathshaala-server/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:root@localhost:5432/pathshaala"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	query := db.New(conn)
	server := server.NewServer(query)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
