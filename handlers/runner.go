package handlers

import (
	"fmt"
	"net"

	pb "gitea.dancheg97.ru/dancheg97/go-pacman/gen/pb/proto/v1"
	"gitea.dancheg97.ru/dancheg97/go-pacman/packages"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	Port int
	*packages.Packager
}

func Run(params *Params) error {
	server := grpc.NewServer(
		getUnaryMiddleware(),
		getStreamMiddleware(),
	)

	handlers := Handlers{
		Packager: params.Packager,
	}
	pb.RegisterAurServiceServer(server, handlers)
	reflection.Register(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(`:%d`, params.Port))
	if err != nil {
		return err
	}

	logrus.Infof(`Grpc server started on port: %d`, params.Port)
	return server.Serve(lis)
}