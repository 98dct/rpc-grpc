package main

import (
	u "5-12-rpc/grpc/proto/user"
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserService struct {
	u.UnimplementedUserServer
}

func (*UserService) GetUser(c context.Context, req *u.GetUserReq) (*u.GetUserResp, error) {
	user, ok := user[req.Id]
	if !ok {
		return nil, errors.New("查无此人")
	}

	return &u.GetUserResp{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	u.RegisterUserServer(s, new(UserService))

	log.Println("服务启动")
	s.Serve(listener)
}
