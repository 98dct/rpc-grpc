package main

import (
	"log"
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

func main() {

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(err)
	}

	defer client.Close()

	req := GetUserReq{Id: "1"}
	resp := GetUserResp{}

	err = client.Call("UserService.GetUser", req, &resp)
	if err != nil {
		panic(err)
	}

	log.Println(resp)

}
