package main

import (
	"context"
	"log"

	"customer-service/proto/gen"

	"github.com/segmentio/kafka-go"
)

type grpcServer struct {
	gen.UnimplementedConsumerServer
	consumer *kafka.Reader
}

func (s *grpcServer) GetLocation(req *gen.GetLocationReq, stream gen.Consumer_GetLocationServer) error {
	for {
		msg, err := s.consumer.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			return err
		}

		var location Location
		location.fromJSON(msg.Value)
		log.Println(location)

		if err := stream.Send(&gen.LocationResp{
			Long: location.Long,
			Lat:  location.Lat,
		}); err != nil {
			log.Println(err)
			return err
		}
	}
}
