package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/bavix/gripmock/protogen/example/multi-package"
	multi_package "github.com/bavix/gripmock/protogen/example/multi-package/bar"
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
	name := "tokopedia"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.Greet(context.Background(), &multi_package.Bar{Name: name})
	if err != nil {
		log.Fatalf("error from grpc: %v", err)
	}

	log.Printf("Greeting: %s", r.Response)
}
