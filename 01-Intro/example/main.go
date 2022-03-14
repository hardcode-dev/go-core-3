package main

import (
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", factorial)
	http.ListenAndServe(":7777", nil)
}

func factorial(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/")
	n, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i
	}
	w.Write([]byte(strconv.Itoa(f)))
}
