package main

import (
	"database/sql"
	"log"

	"github.com/stanlee321/facebook-ads-server/api"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:root@localhost:5432/facebook_ads?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)


func main(){
	conn, err := sql.Open(dbDriver,dbSource)

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)


	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}