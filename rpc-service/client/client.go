package main

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/qiuqiu1999/go-blog/rpc-service/internal/middleware"
	pb "github.com/qiuqiu1999/go-blog/rpc-service/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	ctx := context.Background()

	clientConn, _ := GetClientConn(ctx, "localhost:8001", []grpc.DialOption{grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(middleware.UnaryContextTimeout()),
	)}) //如果想马上连通服务端 []grpc.DialOption{grpc.WithBlock()}
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)

	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Printf("%s", err)
	}
	log.Printf("resp: %v", resp)

}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	/*opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			middleware2.UnaryContextTimeout(),

		),
	))*/
	/*opts = append(opts, grpc.WithStreamInterceptor(
		grpc_middleware.ChainStreamClient(
			middleware2.StreamContextTimeout(),
		),
	))

	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(2),
				grpc_retry.WithCodes(
					codes.Unknown,
					codes.Internal,
					codes.DeadlineExceeded,
				),
			),
		),
	))*/
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}