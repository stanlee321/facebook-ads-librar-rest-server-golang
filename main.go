package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/stanlee321/facebook-ads-server/api"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
	pb "github.com/stanlee321/facebook-ads-server/pkg/api/v1"
	pb_etl "github.com/stanlee321/facebook-ads-server/pkg/etl/api/v1"

	"google.golang.org/grpc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:root@localhost:5432/facebook_ads?sslmode=disable"
	serverAddress = "0.0.0.0:8080"

	gRPCAddress    = "0.0.0.0:4040"
	gRPCETLAddress = "0.0.0.0:50051"
)

func main() {
	var err error
	opts := grpc.WithInsecure()

	// Dial to Service
	connFacebook, err := grpc.Dial(gRPCAddress, opts)
	if err != nil {
		panic(err)
	}

	// Dial to ETL Service
	connFacebookETL, err := grpc.Dial(gRPCETLAddress, opts)
	if err != nil {
		panic(err)
	}

	// Create instance for grpc client
	facebookClient := pb.NewFacebookAdsServiceClient(connFacebook)
	defer connFacebook.Close()

	// Create instance for grpc ETL client
	facebookETLClient := pb_etl.NewFacebookAdsETLServiceClient(connFacebookETL)
	defer connFacebookETL.Close()

	// Open DB Connection
	connDB, err := sql.Open(dbDriver, dbSource)

	// Create store instance
	store := db.NewStore(connDB)

	// Instance of server with params
	server := api.NewServer(store, facebookClient, facebookETLClient)

	// Init Server
	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
