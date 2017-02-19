package ServerManager

import (
	"UserServer"
	"fmt"
	"golang.org/x/net/context"
)

// Define UsernameAvailable with server datatype
func (s *server) CheckUserExists(ctx context.Context, in *userserver.UserQuery) (*userserver.UsernameAvailable, error) {
	var status bool = in.Username == "UNAME"
	fmt.Println("check username")
	fmt.Println(in.Username)
	return &userserver.UsernameAvailable{Exists: status}, nil
}

// Define LoginUser with server
func (s *server) UserLogin(ctx context.Context, in *userserver.LoginRequest) (*userserver.LoginReply, error) {
	fmt.Println("Login")
	fmt.Println(in.Username)
	fmt.Println(in.Password)
	return &userserver.LoginReply{Status: 0}, nil
}

// Define RegisterUser
func (s *server) RegisterUser(ctx context.Context, in *userserver.RegisterRequest) (*userserver.RegisterStatus, error) {
	fmt.Println("Register user")
	fmt.Println(in.Username)
	fmt.Println(in.Email)
	fmt.Println(in.FistName)
	fmt.Println(in.LastName)
	fmt.Println(in.Password)
	
	return &userserver.RegisterStatus{Status: 0}, nil
}