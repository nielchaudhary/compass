package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	listen(":8090")
}

func listen(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "compass server live on 8090!")
	})

	log.Printf("compass server live on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}