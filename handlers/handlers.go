package handlers

import (
	"context"
	"strings"

	pb "gitea.dancheg97.ru/dancheg97/go-pacman/gen/pb/proto/v1"
	"gitea.dancheg97.ru/dancheg97/go-pacman/packages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handlers struct {
	Packager *packages.Packager
}

var ErrUnknown = status.Error(codes.NotFound, `unknown error`)

func (s Handlers) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	out, err := s.Packager.Add(in.Packages)
	if err != nil {
		return nil, err
	}
	return &pb.AddResponse{
		Message: strings.Split(*out, "\n"),
	}, nil
}

func (s Handlers) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	rez, err := s.Packager.Search(in.Pattern)
	if err != nil {
		return nil, err
	}
	return &pb.SearchResponse{
		Packages: rez,
	}, nil
}
