package httpservers

import (
	"fmt"

	"context"
	"net/http"

	pb "dancheg97.ru/dancheg97/ctlpkg/gen/proto/v1"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Params struct {
	HttpPort  int
	GrpcPort  int
	ApiPath   string
	PkgsDir   string
	WebDir    string
	PacmanSvc pb.PacmanServiceServer
}

func RunHttpWrapper(params Params) error {
	ctx := context.Background()
	svmux := runtime.NewServeMux()
	r := mux.NewRouter()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterPacmanServiceHandlerFromEndpoint(ctx, svmux, `localhost:9080`, opts)
	if err != nil {
		return err
	}

	fs1 := http.FileServer(http.Dir(params.WebDir))
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", fs1))

	fs2 := http.FileServer(http.Dir(params.PkgsDir))
	r.PathPrefix("/pkg/").Handler(http.StripPrefix("/pkg/", fs2))

	r.PathPrefix("/").Handler(svmux)

	http.Handle("/", r)
	logrus.Info("server running...")

	return http.ListenAndServe(fmt.Sprintf(`:%d`, params.HttpPort), r)
}
