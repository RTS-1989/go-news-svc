package client

import (
	"fmt"
	"github.com/RTS-1989/go-news-svc/pkg/middleware"
	"github.com/RTS-1989/go-news-svc/pkg/pb/comment"

	"github.com/RTS-1989/go-news-svc/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client comment.CommentServiceClient
}

func InitServiceClient(c *config.Config) comment.CommentServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CommentSvcUrl, grpc.WithInsecure(), middleware.WithClientUnaryInterceptor())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return comment.NewCommentServiceClient(cc)
}
