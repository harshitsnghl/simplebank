package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/harshitsnghl/simplebank/api"
	db "github.com/harshitsnghl/simplebank/db/sqlc"
	"github.com/harshitsnghl/simplebank/util"
	_ "github.com/lib/pq"
)

// Note lib pq driver is used to talk to the database, without this our code would not be able to talk to the database. Hence we did a blind import
// This is also used in main_test.go file

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
