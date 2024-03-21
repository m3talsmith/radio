package main

import (
	"errors"
	"io"
	"log"

	pb "github.com/m3talsmith/radio/server/radio"
)

type service struct {
	pb.UnimplementedRadioAPIServer
	listenerChan chan *listener
}

func (s *service) Station(stream pb.RadioAPI_StationServer) error {
	var currentId string

	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if currentId != "" {
					deleteListener(currentId)
				}
				return nil
			}
			log.Printf("[ERROR] unknown error for listener %s: %v", currentId, err)
			return err
		}

		currentId = req.GetCallerId()

		if l := getListener(currentId); l == nil {
			l = newListener(currentId, &stream)
			setListener(s.listenerChan, l)
		}

		log.Printf("[BROADCASTING] %s from %s to %d listeners", req.GetMessage(), req.GetCallerId(), len(listeners))

		others := otherListeners(currentId)
		res := pb.Broadcast{Message: req.GetMessage(), CallerId: req.GetCallerId()}
		for _, lis := range others {
			if err := (*lis.Stream).Send(&res); err != nil {
				log.Printf("[ERROR] broadcasting to listener %s: %v", lis.Id, err)
				log.Printf("[ERROR DETAILS] %v", (*lis.Stream).Context().Err())
				deleteListener(lis.Id)
			}
		}
	}
}

func newListener(id string, stream *pb.RadioAPI_StationServer) *listener {
	return &listener{Id: id, Stream: stream, Online: true}
}

func getListener(id string) *listener {
	for _, l := range listeners {
		if l.Id == id {
			return l
		}
	}
	return nil
}

func setListener(c chan *listener, l *listener) {
	c <- l
}

func deleteListener(id string) {
	for _, l := range listeners {
		if l.Id == id {
			l.Online = false
		}
	}
}

func otherListeners(id string) []*listener {
	var newListeners []*listener
	for _, l := range listeners {
		if l.Id != id {
			newListeners = append(newListeners, l)
		}
	}
	return newListeners
}
