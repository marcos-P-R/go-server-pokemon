package main

import (
	"fmt"
	"log"
	"module/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var value string = api.Conexao()
	fmt.Fprint(w, value)
}
