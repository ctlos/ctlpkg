package fileserver

import (
	"fmt"
	"net/http"
	"time"
)

type PathPair struct {
	LocalPath  string
	HandlePath string
}

func RunStaticFileServer(pathes []PathPair, port int) {
	for _, pp := range pathes {
		fs := http.FileServer(http.Dir(pp.LocalPath))
		http.Handle(pp.HandlePath, fs)
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(`:%d`, port),
		ReadHeaderTimeout: time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
