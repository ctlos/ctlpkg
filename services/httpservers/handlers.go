package httpservers

import (
	"fmt"
	"time"

	"context"
	"net/http"

	pb "dancheg97.ru/dancheg97/ctlpkg/gen/proto/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Params struct {
	HttpPort          int
	GrpcPort          int
	StaticFileServers []PathPair
}

type PathPair struct {
	LocalPath  string
	HandlePath string
}

func RunHttp(params Params) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterPacmanServiceHandlerFromEndpoint(ctx, mux, "localhost:"+fmt.Sprint(params.GrpcPort), opts)
	if err != nil {
		return err
	}

	http.Handle(`/api`, mux)

	for _, sfs := range params.StaticFileServers {
		fs := http.FileServer(http.Dir(sfs.LocalPath))
		http.Handle(sfs.HandlePath, fs)
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(`:%d`, params.HttpPort),
		ReadHeaderTimeout: time.Second,
	}

	return server.ListenAndServe()
}
