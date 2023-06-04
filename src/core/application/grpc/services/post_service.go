package grpcservice

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/kainguyen/go-scrapper/src/core/application/common/persistence"
	"github.com/kainguyen/go-scrapper/src/core/application/grpc/pb"
	"github.com/kainguyen/go-scrapper/src/core/domain/enums"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/utils"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"io"
	"os"
	"path/filepath"
	"time"
)

type PostServiceServer struct {
	db           *gorm.DB                  `di.scope:"singleton" di.inject:"db"`
	redisService persistence.IRedisService `di.scope:"singleton" di.inject:"redis"`
	logger       *zerolog.Logger           `di.scope:"Singleton" di.inject:"logger"`

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

func (ps *PostServiceServer) GetPostByIds(req *pb.GetPostByIdsRequest, stream pb.PostService_GetPostByIdsServer) error {

	for _, id := range req.GetPostIds() {
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return err
		}

		var post models.Post

		tx := ps.db.First(&post, parsedId)

		if tx.RowsAffected != 0 {
			var pbPost pb.Post

			err := copier.Copy(&pbPost, &post)
			if err != nil {
				return status.Errorf(codes.Internal, fmt.Sprintf("Unexpected Error: %v", err))
			}

			res := &pb.GetPostByIdsResponse{
				Post: &pbPost,
			}
			err = stream.Send(res)

			if err != nil {
				return status.Errorf(codes.Internal, fmt.Sprintf("Unexpected Error: %v", err))
			}
		}
	}

	return nil
}

func WriteFile(dir string, fileName string, buffer *bytes.Buffer) error {
	// Concat dir with current working directory(where main.go lives)
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir = cwd + dir

	// Concat current time to fileName
	current := time.Now()
	fileName = fmt.Sprintf("%v_%s", current.Format(time.RFC3339), fileName)

	file, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (ps *PostServiceServer) UploadPostImage(stream pb.PostService_UploadPostImageServer) error {
	ps.logger.Info().Msg("Receive upload post image")

	req, err := stream.Recv()
	if err != nil {
		ps.logger.Error().Err(err).Msg("Cannot get request information")
		return status.Errorf(codes.Unknown, fmt.Sprintf("Cannot get request information"))
	}

	postId := req.GetInfo().GetPostId()

	id, err := uuid.Parse(postId)
	if err != nil {
		ps.logger.Error().Err(err).Msg("Cannot parse request postId")
		return status.Errorf(codes.Internal, fmt.Sprint(err))
	}

	var post = models.Post{
		Id: id,
	}

	tx := ps.db.Preload("PostImages").Find(&post)
	if tx.RowsAffected == 0 {
		return status.Errorf(codes.NotFound, "Post %v not found", req.GetInfo().GetPostId())
	}

	ps.logger.Info().Msg(fmt.Sprintf("Post found with id: %+v", post.Id))

	imageSize := 0
	imageData := bytes.Buffer{}

	ps.logger.Info().Msg("Waiting to receive post image chunks")

	for {
		contextError := utils.ContextError(stream.Context())
		if contextError != nil {
			ps.logger.Error().Err(contextError).Msg("Context error")
			return contextError
		}

		req, err := stream.Recv()
		if err == io.EOF {
			ps.logger.Warn().Err(err).Msg("No more data")
			break
		}

		if err != nil {
			return status.Errorf(codes.Internal, "Error when receiving stream request: %v", err)
		}

		var chunk = req.GetChunkData()
		imageSize += len(chunk)

		_, err = imageData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "Cannot write data chunk: %v", err)
		}
	}

	fileName := fmt.Sprintf("%v%v", post.Id, req.GetInfo().GetType())

	dir := "/assets/storage"

	err = WriteFile(dir, fileName, &imageData)
	if err != nil {
		return status.Errorf(codes.Internal, "Cannot write image to file: %v", err)
	}

	res := &pb.UploadImageResponse{
		Id:   req.GetInfo().GetPostId(),
		Size: int64(imageSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		ps.logger.Error().Err(err).Msg("Cannot reponse to client")
		return err
	}

	return nil
}
