package server

import (
	"context"
	"encoding/json"
	rapi "github.com/qiuqiu1999/go-blog/rpc-service/pkg/api"
	error2 "github.com/qiuqiu1999/go-blog/rpc-service/pkg/error"
	pb "github.com/qiuqiu1999/go-blog/rpc-service/proto"
)

type TagServer struct {}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := rapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, error2.TogRPCError(error2.ErrorGetTagListFail)
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, error2.TogRPCError(error2.Fail)
	}

	return &tagList, nil
}
