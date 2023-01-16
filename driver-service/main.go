package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"driver-service/proto/gen"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()
	flag.Parse()

	// init depedency
	conn := initKafka("driver-topic", "tcp", os.Getenv("KAFKA_ADDR"), 0)

	// init server
	s := grpc.NewServer()
	gen.RegisterDriverServer(s, &server{
		p: producer{
			conn: conn,
		},
	})

	// setup port
	portInt, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("error parsing PORT from .env: %v", portInt)
	}
	port := flag.Int("port", portInt, "The server port")

	// listen to port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
