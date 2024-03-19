package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/m3talsmith/radio/server/radio"

	"google.golang.org/grpc"
)

var (
	host     string
	grpcPort int
	httpPort int
)

func init() {
	flag.StringVar(&host, "host", "localhost", "Server host")
	flag.IntVar(&grpcPort, "grpc-port", 3001, "GRPC server port")
	flag.IntVar(&httpPort, "http-port", 3000, "HTTP server port")
	flag.Parse()
}

func main() {
	// http section
	// go func() {
	// }()

	// grpc section
	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterRadioAPIServer(server, &service{})
	log.Printf("GRPC listening at %v", conn.Addr())
	log.Fatal(server.Serve(conn))
}
