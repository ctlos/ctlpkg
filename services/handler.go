package services

import (
	"context"
	"fmt"

	"gitea.dancheg97.ru/dancheg97/regen/gen/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handlers struct{}

var ErrUnknown = status.Error(codes.NotFound, `unknown error`)

func (s Handlers) Add(ctx context.Context, in *pb.Package) (*pb.Empty, error) {
	fmt.Println("nan")
	return &pb.Empty{}, nil
}
