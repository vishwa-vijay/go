package main

import (
	"context"
	"fmt"
	"log"
	"net"

	book "github.com/vishwa-vijay/go/excercise/excercise5/proto"
	"google.golang.org/grpc"
)

// Server : is a server
type Server struct {
}

// GetBook : returns books
func (s *Server) GetBook(context context.Context, request *book.GetBookRequest) (*book.GetBookResponse, error) {
	log.Println("Request came to get book", request)
	return &book.GetBookResponse{
		Name: "Vishwa Vijay",
	}, nil
}

func main() {
	log.Println("GRPC server is going to start.")
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "8080"))
	if err != nil {
		log.Println("Error listening", err)
	}
	log.Println("Server started.")

	grpcServer := grpc.NewServer()
	book.RegisterBookServiceServer(grpcServer, &Server{})
	grpcServer.Serve(lis)
}
