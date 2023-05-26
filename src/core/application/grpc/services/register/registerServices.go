package grpcservice

import (
	"github.com/goioc/di"
	"github.com/kainguyen/go-scrapper/src/core/application/grpc/pb"
	grpcservice "github.com/kainguyen/go-scrapper/src/core/application/grpc/services"
	"google.golang.org/grpc"
)

func RegisterServices(server grpc.ServiceRegistrar) {
	// Post service
	postService, err := di.GetInstanceSafe("postServiceServer")
	if err != nil {
		return
	}

	pb.RegisterPostServiceServer(server, postService.(*grpcservice.PostServiceServer))
}
