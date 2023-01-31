package pacman

import (
	"context"
	"fmt"
	"os"

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
	err := s.Helper.Execute("yay -Sy " + in.Packages)
	if err != nil {
		return nil, fmt.Errorf("unable to execute yay command: %w", err)
	}
	err = os.Chmod(s.YayPath, 0777)
	if err != nil {
		return nil, err
	}
	err = os.Chmod(s.PkgPath, 0777)
	if err != nil {
		return nil, err
	}
	entires, err := os.ReadDir(s.YayPath)
	if err != nil {
		return nil, err
	}
	for _, de := range entires {
		if de.IsDir() {
			err := s.Helper.Move(s.YayPath+"/"+de.Name(), s.PkgPath)
			if err != nil {
				return nil, err
			}
		}
	}
	repo := s.PkgPath + "/" + s.RepoName + ".db.tar.gz"
	pkgs := s.PkgPath + "/*.pkg.tar.zst"
	err = s.Helper.Execute("repo-add -n -q " + repo + " " + pkgs)
	if err != nil {
		return nil, err
	}
	return &pb.AddResponse{}, nil
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
	err := s.Helper.Execute("yay -Syu")
	return &pb.UpdateResponse{}, err
}
