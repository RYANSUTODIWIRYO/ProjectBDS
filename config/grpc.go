package config

import (
	//"fmt"
	"log"
	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	)

func KoneksiGrpc()(conn *grpc.ClientConn, er error){
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn, nil
}