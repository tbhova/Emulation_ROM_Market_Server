package ServerManager

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"UserServer"
)

// Data type use to implement login server service
type LoginServer struct{}

func RunLoginServer(s *grpc.Server) {
	userserver.RegisterUserServerServer(s, &LoginServer{})
}

// Define UsernameAvailable with GreetingServer datatype
func (s *LoginServer) CheckUserExists(ctx context.Context, in *userserver.UserQuery) (*userserver.UsernameAvailable, error) {
	var status bool = in.Username == "UNAME"
	fmt.Println("check username")
	fmt.Println(in.Username)
	return &userserver.UsernameAvailable{Exists: status}, nil
}

// Define LoginUser with GreetingServer
func (s *LoginServer) UserLogin(ctx context.Context, in *userserver.LoginRequest) (*userserver.LoginReply, error) {
	fmt.Println("Login")
	fmt.Println(in.Username)
	fmt.Println(in.Password)
	return &userserver.LoginReply{Status: 0}, nil
}

// Define RegisterUser
func (s *LoginServer) RegisterUser(ctx context.Context, in *userserver.RegisterRequest) (*userserver.RegisterStatus, error) {
	fmt.Println("Register user")
	fmt.Println(in.Username)
	fmt.Println(in.Email)
	fmt.Println(in.FistName)
	fmt.Println(in.LastName)
	fmt.Println(in.Password)
	
	return &userserver.RegisterStatus{userserver.RegisterStatus_OK}, nil
}