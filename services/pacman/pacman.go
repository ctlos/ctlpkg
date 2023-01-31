package pacman

import (
	"context"
	"fmt"

	pb "gitea.dancheg97.ru/dancheg97/go-pacman/gen/pb/proto/v1"
	"gitea.dancheg97.ru/dancheg97/go-pacman/src"
)

type Handlers struct {
	Helper   *src.OsHelper
	YayPath  string
	PkgPath  string
	RepoName string
}

func (s Handlers) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	err := s.Helper.Execute("yay -Sy --noconfirm " + in.Packages)
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
