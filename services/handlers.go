package services

import (
	"fmt"
	"net"

	"gitea.dancheg97.ru/dancheg97/regen/gen/pb"
	"gitea.dancheg97.ru/dancheg97/regen/postgres"
	"gitea.dancheg97.ru/dancheg97/regen/services/users"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	Postgres *postgres.Postgres
	Port     int
}

func Run(params *Params) error {
	server := grpc.NewServer(
		getUnaryMiddleware(),
		getStreamMiddleware(),
	)

	handlers := users.Handlers{Postgres: params.Postgres}
	pb.RegisterUserStorageServer(server, handlers)
	reflection.Register(server)

	lis, err := net.Listen("tcp", fmt.Sprintf(`:%d`, params.Port))
	if err != nil {
		return err
	}
	logrus.Infof(`Grpc server started on port: %d`, params.Port)
	return server.Serve(lis)
}
