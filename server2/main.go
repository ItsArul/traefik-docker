package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World From Server 2")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:3001", nil))
}
