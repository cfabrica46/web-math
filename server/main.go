package main

import (
	"log"
)

func main() {
	err := runServer("8080", "8081")
	if err != nil {
		log.Fatal(err)
	}
}
