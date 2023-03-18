package httpservers

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	PkgsDir   string
	WebDir    string
	PacmanSvc pb.PacmanServiceServer
}

func WebForward(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	http.Redirect(w, r, "/web", http.StatusOK)
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
	r.PathPrefix("/web").Handler(http.StripPrefix("/web", fs1))

	fs2 := http.FileServer(http.Dir(params.PkgsDir))
	r.PathPrefix("/pkg").Handler(http.StripPrefix("/pkg", fs2))

	rootPatt := runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{""}, ""))
	svmux.Handle("GET", rootPatt, WebForward)

	r.PathPrefix("/").Handler(svmux)

	http.Handle("/", r)
	logrus.Info("server running...")

	server := http.Server{
		Addr:              fmt.Sprintf(`:%d`, params.HttpPort),
		Handler:           r,
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
	}

	return server.ListenAndServe()
}
