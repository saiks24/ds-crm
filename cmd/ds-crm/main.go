package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/saiks24/ds-crm/internal/app"
	ds_crmpb "github.com/saiks24/ds-crm/pkg/api"
)

func main() {
	ctx := context.Background()

	if err := initCfg(); err != nil {
		log.Fatal(err)
	}

	api := app.NewDsCrm(nil)

	go grpcEndpointServe(viper.GetString("listen_grpc"), api)

	grpcProxyServe(ctx, viper.GetString("listen_http"), viper.GetString("listen_grpc"))
}

func grpcEndpointServe(grpcEndpointAddr string, svc ds_crmpb.DSCrmServer) {
	con, err := net.Listen("tcp", grpcEndpointAddr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	ds_crmpb.RegisterDSCrmServer(server, svc)

	fmt.Println("starting gRPC server at " + grpcEndpointAddr)

	if err := server.Serve(con); err != nil {
		log.Fatal(err)
	}
}

func grpcProxyServe(ctx context.Context, proxyAddr string, grpcEndpointAddr string) {
	grpcConn, err := grpc.Dial(grpcEndpointAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer grpcConn.Close()

	grpcGWMux := runtime.NewServeMux()

	err = ds_crmpb.RegisterDSCrmHandler(ctx, grpcGWMux, grpcConn)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("starting http server at " + proxyAddr)

	log.Fatal(http.ListenAndServe(proxyAddr, grpcGWMux))
}

func initCfg() error {
	pathToConfig := flag.String("config", "", "path to cfg file")

	flag.Parse()

	if pathToConfig == nil || *pathToConfig == "" {
		return errors.New("missing config")
	}

	viper.SetConfigFile(*pathToConfig)

	return viper.ReadInConfig()
}
