package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/d2e-tech/scatter/pb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
    pb.UnimplementedScatterServer
}

func (s *server) Ingest(ctx context.Context, in *pb.IngestRequest) (*pb.IngestReply, error) {
    log.Printf("Received: %v", in.GetObservations())
    return &pb.IngestReply{}, nil
}


func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScatterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
