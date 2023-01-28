package api

import (
	"fmt"
	"net"

	"gitea.dancheg97.ru/dancheg97/go-pacman/gen/pb"
	"gitea.dancheg97.ru/dancheg97/go-pacman/pkg"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	Port int
	*pkg.Packager
}

func Run(params *Params) error {
	server := grpc.NewServer(
		getUnaryMiddleware(),
		getStreamMiddleware(),
	)

	handlers := Handlers{
		Packager: params.Packager,
	}
	pb.RegisterAurServer(server, handlers)
	reflection.Register(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(`:%d`, params.Port))
	if err != nil {
		return err
	}

	logrus.Infof(`Grpc server started on port: %d`, params.Port)
	return server.Serve(lis)
}
