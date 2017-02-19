package ServerManager

import (
	"golang.org/x/net/context"
	"net"
	"log"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	pb "MarketServer"
	"UserServer"
)

var s *grpc.Server

func init() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

const (
	port = ":50051"
)

// server is used to implement server.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// Define UsernameAvailable with server
func (s *server) CheckUser(ctx context.Context, in *userserver.UserQuery) (*userserver.UsernameAvailable, error) {
	var status bool = in.Username == "UNAME"
	return &userserver.UsernameAvailable{Exists: status}, nil
}

//TODO other userserver methods in server



func RunServers() {
	runLoginServer()
	runAvailableDownloads()
	runDownloadServer()
}

func runLoginServer() {
	userserver.RegisterUserServerServer(s, &server{})
}

func runAvailableDownloads() {
	
}

func runDownloadServer() {
	
}