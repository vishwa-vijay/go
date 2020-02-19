package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/jsonpb"
	_struct "github.com/golang/protobuf/ptypes/struct"
	book "github.com/vishwa-vijay/go/excercise/excercise6/proto"
	"google.golang.org/grpc"
)

// Server : is a server
type Server struct {
}

// GetBook : returns books
func (s *Server) GetBook(context context.Context, request *book.GetBookRequest) (*_struct.Struct, error) {
	log.Println("Request came to get book", request)
	var msg string = "{\"Name\":\"John\",\"address\":{\"state\":\"NY\"}}"
	resposne := &_struct.Struct{}
	//err := resposne.XXX_Unmarshal([]byte(msg))
	err := jsonpb.UnmarshalString(msg, resposne)
	return resposne, err
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
