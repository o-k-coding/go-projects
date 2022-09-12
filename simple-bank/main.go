package main

import (
	"context"
	"log"
	"net"
	"net/http"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/okeefem2/simple_bank/api"
	"github.com/okeefem2/simple_bank/config"
	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/gapi"
	"github.com/okeefem2/simple_bank/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// SO a note here, this pattern is more about passing objects needed around
	// the other was more about creating objects that had access to the things needed,
	// then creating receiver functions on those. A more OOP approach.
	c, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn := db.ConnectPostgres(c)
	store := db.NewStore(conn)

	if c.ServerType == "http" {
		runHttpServer(store, c)
	} else if c.ServerType == "grpc" || c.ServerType == "gateway" {
		// This needs to come first because runGrpcServer is a blocking call
		// In a real life I would have a different log prefix or something for both of these so I know which server is doing what.
		// Or maybe would run them on separate processes?
		if c.ServerType == "gateway" {
			go runGatewayServer(store, c)
		}
		runGrpcServer(store, c)

	} else {
		log.Fatalf("server type not supported: %s", c.ServerType)
	}
}

func runHttpServer(store db.Store, conf *config.Config) {
	server, err := api.NewServer(store, *conf)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start http server", err)
	}
}

func runGrpcServer(store db.Store, conf *config.Config) {
	server, err := gapi.NewServer(store, *conf)
	if err != nil {
		log.Fatal("cannot create grpc server", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	// Allows the client to explore what RPCs are availabe on the server, and how to call them
	// documentation for the server basically.
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatal("cannot start grpc listener", err)
	}
	log.Printf("started gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}

func runGatewayServer(store db.Store, conf *config.Config) {
	server, err := gapi.NewServer(store, *conf)
	if err != nil {
		log.Fatal("cannot create grpc server", err)
	}

	grpcMux := runtime.NewServeMux()
	grpcCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = pb.RegisterSimpleBankHandlerServer(grpcCtx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register grpc gateway server", err)
	}

	httpMux := http.NewServeMux()
	httpMux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start grpc gateway listener", err)
	}
	log.Printf("started gRPC gateway server at %s", listener.Addr().String())

	err = http.Serve(listener, httpMux)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
