package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DataLoad(w http.ResponseWriter, r *http.Request) {
	response, err := ioutil.ReadFile("data/data.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	fmt.Fprintf(w, string(response))
}

type Data struct {
	PlayTime      int
	CountdownTime int
	CurrentTime   float32
	Volume        float32
}

func DataSave(w http.ResponseWriter, r *http.Request) {
	var d Data
	err := json.NewDecoder(r.Body).Decode(&d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}

	file, _ := json.MarshalIndent(d, "", " ")
	_ = ioutil.WriteFile("data/data.json", file, 0755)
}
