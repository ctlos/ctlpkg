package server

import (
	"fmt"
	"net/http"
	"time"
)

func RunFileServer(path string, port int) {
	fs := http.FileServer(http.Dir(path))
	http.Handle(`/x86_64/`, fs)

	server := &http.Server{
		Addr:              fmt.Sprintf(`:%d`, port),
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
