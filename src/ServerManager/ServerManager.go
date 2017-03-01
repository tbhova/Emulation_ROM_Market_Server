package ServerManager

import (
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	pb "MarketServer"
)

var s *grpc.Server

func init() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s = grpc.NewServer()
	pb.RegisterGreeterServer(s, &GreetingServer{})
	
	RunLoginServer(s)
	RunUserDownloadServer(s)
	RunAvailableGameServer(s)
	
	// Register reflection service on gRPC GreetingServer.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

const (
	port = ":50051"
)

// Server types used to implement servers.
type GreetingServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *GreetingServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func Run() {
	
}
