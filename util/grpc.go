package util

import (
	"log"

	"google.golang.org/grpc"
)

// Dial grpc server with apm middleware
func Dial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", addr, err)
	}
	return conn
}
