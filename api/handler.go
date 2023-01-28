package api

import (
	"context"
	"fmt"

	"gitea.dancheg97.ru/dancheg97/go-pacman/gen/pb"
	"gitea.dancheg97.ru/dancheg97/go-pacman/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handlers struct {
	Packager *pkg.Packager
}

var ErrUnknown = status.Error(codes.NotFound, `unknown error`)

func (s Handlers) Add(ctx context.Context, in *pb.Package) (*pb.Empty, error) {
	fmt.Println("nan")
	return &pb.Empty{}, nil
}
