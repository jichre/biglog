package main

import (
	"bitbucket.org/biglog/handler"
	"net/http"
	"sync"
)

func main() {
	doneGroup := new(sync.WaitGroup)
	doneGroup.Add(1)

	go HttpServer()

	doneGroup.Wait()
}

func HttpServer() {
	http.HandleFunc("/api/postLog", handler.RevieveHandler)
	http.ListenAndServe(":80", nil)
}
