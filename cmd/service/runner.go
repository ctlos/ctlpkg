package service

import (
	"fmt"
	"net/http"
	"time"

	pb "dancheg97.ru/dancheg97/ctlpkg/cmd/generated/proto/v1"
	"dancheg97.ru/dancheg97/ctlpkg/cmd/utils"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Params struct {
	Port     int
	PkgPath  string
	YayPath  string
	WebPath  string
	RepoName string
	Packager *utils.OsHelper
}

func Run(params *Params) error {
	mux := http.NewServeMux()

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_recovery.UnaryServerInterceptor(),
			getUnaryLogger(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(),
			getStreamLogger(),
		),
	)
	pb.RegisterPacmanServiceServer(grpcServer, Handlers{})

	appfs := http.FileServer(http.Dir(params.WebPath))
	mux.Handle("/", http.StripPrefix("/", appfs))

	pkgfs := http.FileServer(http.Dir(params.PkgPath))
	mux.Handle("/pkg/", http.StripPrefix("/pkg/", pkgfs))

	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	for _, path := range grpcweb.ListGRPCResources(grpcServer) {
		mux.Handle(path, wrappedGrpc)
	}

	server := http.Server{
		Addr:              ":" + fmt.Sprint(params.Port),
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
		Handler:           mux,
	}

	logrus.Info("server running on port: " + fmt.Sprint(params.Port))
	return server.ListenAndServe()
}
