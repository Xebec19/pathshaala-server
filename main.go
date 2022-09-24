package main

import (
	"database/sql"
	"log"

	"github.com/Xebec19/pathshaala-server/api"
	db "github.com/Xebec19/pathshaala-server/db/sqlc"
	"github.com/Xebec19/pathshaala-server/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	query := db.New(conn)
	server := api.NewServer(query)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
