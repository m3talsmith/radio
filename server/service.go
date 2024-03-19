package main

import (
	"errors"
	"io"
	"log"
	"sync"

	pb "github.com/m3talsmith/radio/server/radio"
)

var (
	mut       sync.Mutex
	listeners []pb.RadioAPI_StationServer
)

type service struct {
	pb.UnimplementedRadioAPIServer
}

func (s service) Station(stream pb.RadioAPI_StationServer) error {
	mut.Lock()
	listeners = append(listeners, stream)
	mut.Unlock()

	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		log.Printf("broadcasting %s to %d listeners", req.Name, len(listeners))
		res := pb.Broadcast{Name: req.Name}
		for _, l := range listeners {
			if err := l.Send(&res); err != nil {
				log.Printf("error broadcasting: %v", err)
			}
		}
	}
}
