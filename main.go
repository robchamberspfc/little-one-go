package main

import (
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./main")))
	//http.ListenAndServe(":3000", nil)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
