package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World Server 1")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))

}
