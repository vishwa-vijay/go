Command to generate .pb.go file
```
protoc --go_out=plugins=grpc:. --proto_path=/home/vishwa/go/src:. book.proto
```