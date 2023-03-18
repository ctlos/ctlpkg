package httpservers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
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
	ApiPath   string
	PkgsDir   string
	WebDir    string
	PacmanSvc pb.PacmanServiceServer
}

func WebForward(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	urlStr := "/web"
	if u, err := url.Parse(urlStr); err == nil {
		oldpath := r.URL.Path
		if oldpath == "" {
			oldpath = "/"
		}
		if u.Scheme == "" {
			if urlStr == "" || urlStr[0] != '/' {
				olddir, _ := path.Split(oldpath)
				urlStr = olddir + urlStr
			}

			var query string
			if i := strings.Index(urlStr, "?"); i != -1 {
				urlStr, query = urlStr[:i], urlStr[i:]
			}

			trailing := strings.HasSuffix(urlStr, "/")
			urlStr = path.Clean(urlStr)
			if trailing && !strings.HasSuffix(urlStr, "/") {
				urlStr += "/"
			}
			urlStr += query
		}
	}

	w.Header().Set("Location", urlStr)
	w.WriteHeader(200)

	if r.Method == "GET" {
		note := "<a href=\"" + urlStr + "\"> redirected </a>.\n"
		fmt.Fprintln(w, note)
	}
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
