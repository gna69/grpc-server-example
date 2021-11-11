package addr

import (
	"../api"
	"context"
)

type GRPCServer struct {}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {
	panic("implement me")
}

func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./addr.proto