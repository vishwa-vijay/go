package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	book "github.com/vishwa-vijay/go/excercise/excercise5/proto"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8080", opts...)

	if err != nil {
		log.Printf("failure while dialing: %v", err)
	}
	defer conn.Close()

	client := book.NewBookServiceClient(conn)
	response, err := client.GetBook(context.Background(), &book.GetBookRequest{
		Isbn: 123,
	})

	data, _ := json.Marshal(response)
	fmt.Println(string(data))

}
