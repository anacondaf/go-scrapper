package grpcservice

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/application/grpc/pb"
	"github.com/kainguyen/go-scrapper/src/core/domain/enums"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type PostServiceServer struct {
	db           *gorm.DB                  `di.scope:"singleton" di.inject:"db"`
	redisService persistence.IRedisService `di.scope:"singleton" di.inject:"redis"`
	pb.UnimplementedPostServiceServer
}

func (ps *PostServiceServer) GetPosts(ctx context.Context, _ *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	var posts []models.Post

	_, err := ps.redisService.Get(ctx, enums.POST_KEY, true, &posts)
	if err != nil {
		if err != redis.Nil {
			return nil, status.Errorf(codes.Internal, fmt.Sprint(err))
		}

		ps.db.Preload("PostImages").Find(&posts)
		ps.redisService.Set(ctx, enums.POST_KEY, posts, 0)
	}

	var pbPost []*pb.Post

	err = copier.Copy(&pbPost, &posts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error When Copying Slice: %v", err))
	}

	return &pb.GetPostsResponse{Posts: pbPost}, nil
}
