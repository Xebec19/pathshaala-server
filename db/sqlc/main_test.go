package db

import "database/sql"

config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M){
	var err error
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(config.db)
	testDB, err = sqlOpen(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:",err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}