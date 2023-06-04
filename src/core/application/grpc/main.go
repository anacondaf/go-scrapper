package main

import (
	"flag"
	grpcClient "github.com/kainguyen/go-scrapper/src/core/application/grpc/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func uploadImageHandler(client *grpcClient.PostClient) {
	postId := "a19ab3cf-fb28-460b-97e3-3c98392e5b56"

	client.UploadPostImage(postId, "../../../../assets/macbook.jpeg")
}

func startGRPCClient() {
	grpcServerAddress := flag.String("address", "localhost:8000", "grpc server address")
	flag.Parse()

	log.Printf("Dial server %s", *grpcServerAddress)

	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(*grpcServerAddress, opts...)
	if err != nil {
		log.Fatal("Cannot dial to grpc server: ", err)
	}

	defer conn.Close()

	postClient := grpcClient.NewPostClient(conn)
	uploadImageHandler(postClient)
}

func main() {
	startGRPCClient()
}
