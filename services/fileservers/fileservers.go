package fileservers

import (
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Params struct {
	Listener net.Listener
	PkgsDir  string
	WebDir   string
}

func RunHttpWrapper(params Params) error {
	mux := http.NewServeMux()

	appfs := http.FileServer(http.Dir(params.WebDir))
	mux.Handle("/", http.StripPrefix("/", appfs))

	pkgfs := http.FileServer(http.Dir(params.PkgsDir))
	mux.Handle("/pkg/", http.StripPrefix("/pkg/", pkgfs))

	server := http.Server{
		Handler:           mux,
		ReadTimeout:       time.Minute,
		ReadHeaderTimeout: time.Minute,
		WriteTimeout:      time.Minute,
		IdleTimeout:       time.Minute,
	}

	logrus.Info("HTTP server running...")
	return server.Serve(params.Listener)
}
