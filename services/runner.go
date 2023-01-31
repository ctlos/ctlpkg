package services

import (
	"fmt"
	"net"

	pb "gitea.dancheg97.ru/dancheg97/go-pacman/gen/pb/proto/v1"
	"gitea.dancheg97.ru/dancheg97/go-pacman/services/fileserver"
	"gitea.dancheg97.ru/dancheg97/go-pacman/services/pacman"
	"gitea.dancheg97.ru/dancheg97/go-pacman/src"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	FilePort int
	GrpcPort int
	Packager *src.Packager
}

func Run(params *Params) error {
	go fileserver.RunFileServer(
		params.Packager.PacmanCacheDir,
		viper.GetInt(`file-port`),
	)

	server := grpc.NewServer(
		getUnaryMiddleware(),
		getStreamMiddleware(),
	)

	handlers := pacman.Handlers{
		Packager: params.Packager,
	}
	pb.RegisterPacmanServiceServer(server, handlers)
	reflection.Register(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(`:%d`, params.GrpcPort))
	if err != nil {
		return err
	}

	logrus.Infof(`Grpc server started on port: %d`, params.GrpcPort)
	return server.Serve(lis)
}
