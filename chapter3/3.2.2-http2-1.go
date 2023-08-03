package main

import (
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"time"
)

const idleTimeout = 5 * time.Minute
const activeTimeout = 10 * time.Minute

func main() {
	var srv http.Server
	srv.Addr = ":8972"
	http.HandleFunc("/hi2", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello http2"))
	})
	http2.ConfigureServer(&srv, &http2.Server{})
	go func() {
		log.Fatal(srv.ListenAndServeTLS("chapter3/server.crt", "chapter3/server.key"))
	}()
	select {}
}
