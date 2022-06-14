package main

import (
	"context"
	"fmt"
	pb "github.com/qiuqiu1999/go-blog/rpc-service/proto"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	clientConn, _ := GetClientConn(ctx, "localhost:8001", nil) //如果想马上连通服务端 []grpc.DialOption{grpc.WithBlock()}
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)

	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	fmt.Printf("resp: %v", resp)

}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}