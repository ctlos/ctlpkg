package fileservers

import (
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Params struct {
	Listener   net.Listener
	GrpcServer *grpc.Server
	PkgsDir    string
	WebDir     string
}

func RunHttpWrapper(params Params) error {
	mux := http.NewServeMux()

	appfs := http.FileServer(http.Dir(params.WebDir))
	mux.Handle("/web/", http.StripPrefix("/", appfs))

	pkgfs := http.FileServer(http.Dir(params.PkgsDir))
	mux.Handle("/pkg/", http.StripPrefix("/pkg/", pkgfs))

	wrappedGrpc := grpcweb.WrapServer(params.GrpcServer)
	mux.Handle("/", wrappedGrpc)

	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if wrappedGrpc.IsGrpcWebRequest(r) {
				wrappedGrpc.ServeHTTP(w, r)
				return
			}
			http.DefaultServeMux.ServeHTTP(w, r)
		}),
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
	}

	logrus.Info("HTTP server running...")
	return server.Serve(params.Listener)
}
