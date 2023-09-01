package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bavix/gripmock/protogen/example/stub-subfolders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//nolint:gomnd
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Set up a connection to the server.
	conn, err := grpc.DialContext(ctx, "localhost:4770", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGripmockClient(conn)

	// Contact the server and print out its response.
	r, err := c.SayHello(context.Background(), &pb.Request{Name: "tokopedia"})
	if err != nil {
		log.Fatalf("error from grpc: %v", err)
	}
	log.Printf("Greeting: %s (return code %d)", r.Message, r.ReturnCode)

	r, err = c.SayHello(context.Background(), &pb.Request{Name: "subbavix"})
	if err != nil {
		log.Fatalf("error from grpc: %v", err)
	}
	log.Printf("Greeting: %s (return code %d)", r.Message, r.ReturnCode)
}
