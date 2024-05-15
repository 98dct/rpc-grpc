package main

import (
	u "5-12-rpc/grpc/proto/user"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := u.NewUserClient(conn)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	userResp, err := client.GetUser(ctx, &u.GetUserReq{Id: "1"})
	if err != nil {
		panic(err)
	}
	log.Println(userResp)
}
