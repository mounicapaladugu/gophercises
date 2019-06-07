package main

import (
	"fmt"
	"net/http"
)

func main() {
	//route handler
	http.HandleFunc("/", hello)

	//start server
	fmt.Println("Starting the server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}
