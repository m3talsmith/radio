package main

import (
	"flag"
	"fmt"
	pb "github.com/m3talsmith/radio/server/radio"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type listener struct {
	Id     string
	Stream *pb.RadioAPI_StationServer
	Online bool
}

var (
	host      string
	grpcPort  int
	httpPort  int
	listeners []*listener
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

	// Listener channel handler
	listenChan := make(chan *listener)
	go func() {
		for {
			l := <-listenChan
			listeners = append(listeners, l)
			log.Printf("[NOTICE] %d listeners online", len(listeners))
		}
	}()

	// Listener Healthcheck
	go func() {
		for {
			if len(listeners) > 0 {
				log.Printf("[HEALTHCHECK] verifying status of %d listeners", len(listeners))
				var newListeners []*listener
				for _, l := range listeners {
					if !l.Online || (*l.Stream).Context().Err() != nil {
						log.Printf("[HEALTHCHECK] listener %s OFFLINE", l.Id)
						continue
					}
					newListeners = append(newListeners, l)
				}
				listeners = newListeners
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}()

	// grpc section
	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, grpcPort))
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterRadioAPIServer(server, &service{listenerChan: listenChan})
	log.Printf("GRPC listening at %v", conn.Addr())
	log.Fatal(server.Serve(conn))
}
