package main

import (
	"api/src/router.go"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando API!")
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(":5000", r))
}
