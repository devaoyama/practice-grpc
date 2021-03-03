package main

import (
	"context"
	"fmt"
	"log"
	"practice-grpc/pb/hello"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()
	client := hello.NewHelloClient(con)
	message := &hello.HelloMessage{Name: "山田太郎"}
	res, _ := client.Hello(context.TODO(), message)
	fmt.Println(res.Msg)
}
