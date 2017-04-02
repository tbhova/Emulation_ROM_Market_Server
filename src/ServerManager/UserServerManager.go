package ServerManager

import (
	"fmt"
	"golang.org/x/net/context"
	"UserServer"
	"strings"
	"database/sql"
	"log"
)

// Data type use to implement login server service
type LoginServer struct{}

func RunLoginServer() {
	UserServer.RegisterUserServerServer(s, &LoginServer{})
}

// Define UsernameAvailable
func (s *LoginServer) CheckUserExists(ctx context.Context, in *UserServer.UserQuery) (*UserServer.UsernameAvailable, error) {
	rows, err := db.Query("SELECT USERNAME FROM USERS WHERE USERNAME = ? OR EMAIL = ?", in.Username, in.Email);
	if err != nil {
		log.Fatal(err)
		return &UserServer.UsernameAvailable{Exists: true}, err
	}
	
	for rows.Next() {
		var username string
		err := rows.Scan(username)
		if err != nil {
			log.Fatal(err)
			return &UserServer.UsernameAvailable{Exists: true}, err
		}
		if len(username) > 0 {
			return &UserServer.UsernameAvailable{Exists: true}, err
		}
	}
	return &UserServer.UsernameAvailable{Exists: false}, nil
}

// Define LoginUser
func (s *LoginServer) UserLogin(ctx context.Context, in *UserServer.LoginRequest) (*UserServer.LoginReply, error) {
	var reply *UserServer.LoginReply
	reply.Id = ""
	
	userExists, err := s.CheckUserExists(ctx, &UserServer.UserQuery{Username: in.Username})
	if err != nil {
		log.Fatal(err)
		reply.Status = UserServer.LoginReply_NONEXISTANT_USERNAME
		return reply, err
	}
	
	if !userExists.Exists {
		reply.Status = UserServer.LoginReply_NONEXISTANT_USERNAME
		return reply, nil
	}
	
	var userId string
	err = db.QueryRow("SELECT ID FROM USER WHERE USER = ? AND PASSWORD = ?").Scan(userId)
	
	switch {
	case err == sql.ErrNoRows:
		reply.Status = UserServer.LoginReply_BAD_PASSWORD
		return reply, nil
	case err != nil:
		log.Fatal(err)
		reply.Status = UserServer.LoginReply_BAD_PASSWORD
	}
	
	reply.Status = UserServer.LoginReply_OK
	reply.Id = userId
	return reply, nil
}

// Define RegisterUser
func (s *LoginServer) RegisterUser(ctx context.Context, in *UserServer.RegisterRequest) (*UserServer.RegisterStatus, error) {
	var reply *UserServer.RegisterStatus
	
	userExists, err := s.CheckUserExists(ctx, &UserServer.UserQuery{Username: in.Username, Email: in.Email})
	if err != nil {
		log.Fatal(err)
		reply.Error = UserServer.RegisterStatus_TAKEN_USERNAME
		return reply, err
	}
	
	if !userExists.Exists {
		reply.Error = UserServer.RegisterStatus_TAKEN_USERNAME
		return reply, nil
	}
	
	var userId string
	err = db.QueryRow("INSERT INTO USER (FIRST_NAME, LAST_NAME, EMAIL, USERNAME, PASSWORD) VALUES (?, ?, ?, ?, ?) RETURNING ID",
		in.FistName, in.LastName, in.Email, in.Username, in.Password,).Scan(userId)
	
	if err != nil {
		log.Fatal(err)
		reply.Error = UserServer.RegisterStatus_TAKEN_USERNAME
		return reply, err
	}
	
	reply.Status = UserServer.RegisterStatus_OK
	reply.UserId = userId
	return reply, nil
}