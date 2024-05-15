package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

type GetUserReq struct {
	Id string `json:"id"`
}

type GetUserResp struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type UserService struct {
}

func (*UserService) GetUser(req GetUserReq, resp *GetUserResp) error {
	u, ok := user[req.Id]
	if ok {
		*resp = GetUserResp{
			Id:    u.Id,
			Name:  u.Name,
			Phone: u.Phone,
		}
		return nil
	}

	return errors.New("查无此人")
}

func main() {

	userService := new(UserService)

	rpc.Register(userService)

	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	log.Println("服务启动成功！")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		//封装了读写拆包等
		go rpc.ServeConn(conn)
	}
}
