package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func rolldice(writter http.ResponseWriter, request *http.Request) {
	roll := 1 + rand.Intn(6)

	resp := strconv.Itoa(roll) + "\n"
	if _, err := io.WriteString(writter, resp); err != nil {
		log.Println("Error writing response: %v\n", err)
	}
}
