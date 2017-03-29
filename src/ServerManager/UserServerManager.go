package ServerManager

import (
	"fmt"
	"golang.org/x/net/context"
	"UserServer"
)

// Data type use to implement login server service
type LoginServer struct{}

func RunLoginServer() {
	UserServer.RegisterUserServerServer(s, &LoginServer{})
}

// Define UsernameAvailable with GreetingServer datatype
func (s *LoginServer) CheckUserExists(ctx context.Context, in *UserServer.UserQuery) (*UserServer.UsernameAvailable, error) {
	
	return &UserServer.UsernameAvailable{Exists: status}, nil
}

// Define LoginUser with GreetingServer
func (s *LoginServer) UserLogin(ctx context.Context, in *UserServer.LoginRequest) (*UserServer.LoginReply, error) {
	fmt.Println("Login")
	fmt.Println(in.Username)
	fmt.Println(in.Password)
	return &UserServer.LoginReply{Status: 0}, nil
}

// Define RegisterUser
func (s *LoginServer) RegisterUser(ctx context.Context, in *UserServer.RegisterRequest) (*UserServer.RegisterStatus, error) {
	fmt.Println("Register user")
	fmt.Println(in.Username)
	fmt.Println(in.Email)
	fmt.Println(in.FistName)
	fmt.Println(in.LastName)
	fmt.Println(in.Password)
	
	return &UserServer.RegisterStatus{UserServer.RegisterStatus_OK}, nil
}