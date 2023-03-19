package pacman

import (
	"context"
	"fmt"
	"strings"

	pb "dancheg97.ru/dancheg97/ctlpkg/gen/proto/v1"
	"dancheg97.ru/dancheg97/ctlpkg/src"
)

type Handlers struct {
	Helper   *src.OsHelper
	YayPath  string
	PkgPath  string
	RepoName string
}

// Describe implements pb.PacmanServiceServer
func (s Handlers) Describe(ctx context.Context, in *pb.DescribeRequest) (*pb.DescribeResponse, error) {
	info, err := s.Helper.Call(`yay -Qi ` + in.Package)
	if err != nil {
		return nil, err
	}
	return &pb.DescribeResponse{
		Fields: s.Helper.ParsePkgInfo(info),
	}, nil
}

// Stats implements pb.PacmanServiceServer
func (Handlers) Stats(context.Context, *pb.StatsRequest) (*pb.StatsResponse, error) {
	panic("unimplemented")
}

func (s Handlers) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	err := s.Helper.Execute("yay -Sy --noconfirm " + strings.Join(in.Packages, ` `))
	if err != nil {
		return nil, fmt.Errorf("unable to execute yay command: %w", err)
	}
	err = s.Helper.FormDb(s.YayPath, s.PkgPath, s.RepoName)
	return &pb.AddResponse{}, err
}

func (s Handlers) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchResponse, error) {
	rez, err := s.Helper.Search(s.PkgPath, in.Pattern)
	if err != nil {
		return nil, err
	}
	return &pb.SearchResponse{
		Packages: rez,
	}, nil
}

func (s Handlers) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	err := s.Helper.Execute("yay -Syu --noconfirm ")
	return &pb.UpdateResponse{}, err
}
