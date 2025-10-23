package main

import (
	"log"

	"github.com/cobbles/ant-simulator/server"
)

func main() {
	srv := server.New()
	log.Println("Starting server on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
