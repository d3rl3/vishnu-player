package main

import (
	"net/http"
	// "github.com/maxence-charriere/go-app/v8/pkg/app"
)

const (
	filePrefix = "/music/"
	root       = "./music"
)

func main() {
	http.HandleFunc("/", PlayerMainFrame)
	http.HandleFunc("/data/load", DataLoad)
	http.HandleFunc("/data/save", DataSave)
	http.HandleFunc(filePrefix, File)
	http.HandleFunc("/service_worker.js", ServiceWorker)

	http.ListenAndServe(":8080", nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
