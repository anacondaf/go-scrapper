package grpc_service

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/kainguyen/go-scrapper/src/core/application/grpc/pb"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type PostServiceServer struct {
	db *gorm.DB
	pb.UnimplementedPostServiceServer
}

func NewPostServiceServer(db *gorm.DB) *PostServiceServer {
	return &PostServiceServer{db: db}
}

func (postService *PostServiceServer) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	var posts []models.Post

	postService.db.Preload("PostImages").Find(&posts)

	var pbPost []*pb.Post

	err := copier.Copy(&pbPost, &posts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error When Copying Slice: %v", err))
	}

	return &pb.GetPostsResponse{Posts: pbPost}, nil
}
