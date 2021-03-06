package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-dev-v3/14-RPC/2-gRPC/books_proto"
)

type Bookinist struct {
	Data []pb.Book

	// Композиция интерфейса.
	pb.BookinistServer
}

func (b *Bookinist) Books(_ *pb.Empty, stream pb.Bookinist_BooksServer) error {
	for _, item := range b.Data {
		stream.Send(&item)
	}
	return nil
}

func (b *Bookinist) AddBook(_ context.Context, book *pb.Book) (*pb.Empty, error) {
	b.Data = append(b.Data, *book)
	return nil, nil
}

func main() {
	srv := Bookinist{}
	srv.Data = append(srv.Data, pb.Book{Id: 2, Title: "The Go Programming Language"}, pb.Book{Id: 1, Title: "1984"})

	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookinistServer(grpcServer, &srv)
	grpcServer.Serve(lis)
}
