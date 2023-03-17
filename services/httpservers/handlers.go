package httpservers

import (
	"fmt"

	"context"
	"net/http"

	pb "dancheg97.ru/dancheg97/ctlpkg/gen/proto/v1"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
)

type Params struct {
	HttpPort  int
	GrpcPort  int
	ApiPath   string
	ServeDir  string
	PacmanSvc pb.PacmanServiceServer
}

func RunHttpWrapper(params Params) error {
	ctx := context.Background()
	svmux := runtime.NewServeMux()
	r := mux.NewRouter()

	err := pb.RegisterPacmanServiceHandlerServer(ctx, svmux, params.PacmanSvc)
	if err != nil {
		return err
	}

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(params.ServeDir))))
	r.PathPrefix("/api/").Handler(svmux)
	http.Handle("/", r)

	logrus.Info("running server")

	return http.ListenAndServe(fmt.Sprintf(`:%d`, params.HttpPort), r)
}
