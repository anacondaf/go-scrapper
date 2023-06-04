package grpcClient

import (
	"bufio"
	"context"
	"github.com/kainguyen/go-scrapper/src/core/application/grpc/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// PostClient is the client connect to PostServiceServer
type PostClient struct {
	service pb.PostServiceClient
}

func NewPostClient(cc grpc.ClientConnInterface) *PostClient {
	service := pb.NewPostServiceClient(cc)

	return &PostClient{
		service: service,
	}
}

func (postClient *PostClient) UploadPostImage(postId string, imagePath string) {
	log.Println("Send upload post image request")

	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("Cannot open file:  ", err)
	}

	defer file.Close()

	chunkSize := 1024
	reader := bufio.NewReader(file)
	buffer := make([]byte, chunkSize)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()

	stream, err := postClient.service.UploadPostImage(ctx)
	if err != nil {
		log.Fatal(err)
	}

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				PostId: postId,
				Type:   filepath.Ext(imagePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		log.Fatal("Cannot send image info to server: ", stream.RecvMsg(nil))
	}

	for {
		numOfBytes, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		req = &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:numOfBytes],
			},
		}

		err = stream.Send(req)
		if err != nil {
			log.Fatal("Error when send image chunk", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("%v bytes", res.GetSize())
}
