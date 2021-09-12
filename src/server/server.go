package main

import (
	"context"
	"log"
	"net"
	pb "github.com/Mdaiki0730/grpc-learning"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloRe› ›ply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
  addr := ":50051"
  lis, err := net.Listen("tcp", addr)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, &server{})
  log.Printf("gRPC server listening on " + addr)
  if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
