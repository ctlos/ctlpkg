package fileserver

import (
	"fmt"
	"net/http"
	"time"
)

func RunFileServer(path string, port int) {
	fs := http.FileServer(http.Dir(path))
	http.Handle(`/`, fs)

	server := &http.Server{
		Addr:              fmt.Sprintf(`:%d`, port),
		ReadHeaderTimeout: time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
