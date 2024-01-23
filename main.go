package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web services from GO")
	})

	http.HandleFunc(("/home"), func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./test.html")
	})
	http.ListenAndServe(":3000", nil)
}
