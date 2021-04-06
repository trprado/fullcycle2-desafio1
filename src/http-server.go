package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func index(w http.ResponseWriter, r *http.Request) {
	f, err := ioutil.ReadFile("./index.html")
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(f))
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}
