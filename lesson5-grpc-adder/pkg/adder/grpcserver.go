package adder

import (
    "context"
    "gopher-school/lesson5-grpc-adder/pkg/api"
)

type GRPCServer struct {
    api.UnimplementedAdderServer
}

func (g *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
    return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
