package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name  string
	IsDir bool
	Mode  os.FileMode
}

const (
	filePrefix = "/music/"
	root       = "./music"
)

func main() {
	http.HandleFunc("/", playerMainFrame)
	http.HandleFunc(filePrefix, file)
	http.HandleFunc("/data", data)
	http.ListenAndServe(":8080", nil)
}

type Data struct {
	PlayTime      int
	CountdownTime int
	CurrentTime   int
	Volume        float32
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func data(w http.ResponseWriter, r *http.Request) {
	fmt.Print("called")

	var d Data

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	file, _ := json.MarshalIndent(d, "", " ")
	fmt.Printf("%+v", d)
	_ = ioutil.WriteFile("data/data.json", file, 0644)
}

func playerMainFrame(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./player.html")
}

func file(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(root, r.URL.Path[len(filePrefix):])
	// fmt.Println(path)
	stat, err := os.Stat(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if stat.IsDir() {
		serveDir(w, r, path)
		return
	}
	http.ServeFile(w, r, path)
}

func serveDir(w http.ResponseWriter, r *http.Request, path string) {
	defer func() {
		if err, ok := recover().(error); ok {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	files, err := file.Readdir(-1)
	if err != nil {
		panic(err)
	}

	fileinfos := make([]FileInfo, len(files), len(files))

	for i := range files {
		fileinfos[i].Name = files[i].Name()
		fileinfos[i].IsDir = files[i].IsDir()
		fileinfos[i].Mode = files[i].Mode()
	}

	j := json.NewEncoder(w)

	if err := j.Encode(&fileinfos); err != nil {
		panic(err)
	}
}
