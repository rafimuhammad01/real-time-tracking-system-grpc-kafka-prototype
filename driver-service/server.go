package main

import (
	"context"

	"driver-service/proto/gen"
)

type server struct {
	gen.UnimplementedDriverServer
	p producer
}

func (s *server) SendLocation(ctx context.Context, location *gen.LocationReq) (*gen.EmptyOKResponse, error) {
	loc := Location{
		Long: location.Long,
		Lat:  location.Lat,
	}

	err := s.p.SendDriverLocation(loc)
	if err != nil {
		return nil, err
	}

	return &gen.EmptyOKResponse{
		Message: "OK",
	}, nil
}
