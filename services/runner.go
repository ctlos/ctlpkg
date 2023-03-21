package services

import (
	"fmt"
	"net"

	pb "dancheg97.ru/dancheg97/ctlpkg/gen/go/proto/v1"
	"dancheg97.ru/dancheg97/ctlpkg/services/fileservers"
	"dancheg97.ru/dancheg97/ctlpkg/services/pacman"
	"dancheg97.ru/dancheg97/ctlpkg/src"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Params struct {
	Port     int
	PkgPath  string
	YayPath  string
	WebPath  string
	RepoName string
	Packager *src.OsHelper
}

func Run(params *Params) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(`:%d`, params.Port))
	if err != nil {
		return err
	}
	logrus.Info("Listening to port ", params.Port)

	m := cmux.New(lis)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := m.Match(cmux.Any())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(params, grpcListener) })
	g.Go(func() error { return httpServe(params, httpListener) })
	g.Go(func() error { return m.Serve() })

	return g.Wait()
}

func grpcServe(params *Params, lis net.Listener) error {
	server := grpc.NewServer(
		getUnaryMiddleware(),
		getStreamMiddleware(),
	)

	pacmanService := pacman.Handlers{
		Helper:   params.Packager,
		YayPath:  params.YayPath,
		PkgPath:  params.PkgPath,
		RepoName: params.RepoName,
	}
	pb.RegisterPacmanServiceServer(server, pacmanService)
	reflection.Register(server)

	logrus.Infof(`gRPC server running...`)
	return server.Serve(lis)
}

func httpServe(params *Params, lis net.Listener) error {
	return fileservers.RunHttpWrapper(fileservers.Params{
		Listener: lis,
		PkgsDir:  params.PkgPath,
		WebDir:   params.WebPath,
	})
}
