package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./main")))
	http.ListenAndServe(":3000", nil)

	http.Handle("/", http.FileServer(http.Dir("./main")))
	http.ListenAndServe("web1", nil)
}
