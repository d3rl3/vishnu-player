package main

import (
	"net/http"
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

	http.ListenAndServe(":8080", nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
