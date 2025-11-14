package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nielchaudhary/compass/internal/config"
	constants "github.com/nielchaudhary/compass/pkg/constants"
)

func main() {
	listen(":8090")
}

func listen(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "compass server live on 8090!")
	})

	env := config.GetEnv("APP_ENV", string(constants.Development))
	log.Println("SERVER ENVIRONMENT:", env)

	log.Printf("compass server live on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
