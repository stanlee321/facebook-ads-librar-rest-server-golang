package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/stanlee321/facebook-ads-server/cmd/server/api"
	db "github.com/stanlee321/facebook-ads-server/db/sqlc"
	pb "github.com/stanlee321/facebook-ads-server/pkg/ads/api/v1"
	pb_etl "github.com/stanlee321/facebook-ads-server/pkg/etl/api/v1"
	"github.com/stanlee321/facebook-ads-server/pkg/util"
	"google.golang.org/grpc"
)

func main() {
	var err error

	// load configs
	conf, err := util.LoadConfig(".", "PROD")

	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	opts := grpc.WithInsecure()

	// Dial to Service
	connFacebook, err := grpc.Dial(conf.GRPCAddress, opts)
	if err != nil {
		panic(err)
	}

	// Dial to ETL Service
	connFacebookETL, err := grpc.Dial(conf.GRPCETLAddress, opts)
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
	connDB, err := sql.Open(conf.DBDriver, conf.DBSource)

	// Create store instance
	store := db.NewStore(connDB)

	// Instance of server with params
	server := api.NewServer(store, facebookClient, facebookETLClient)

	// Init Server
	err = server.Start(conf.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
