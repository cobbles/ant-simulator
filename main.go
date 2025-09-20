package main

import (
	"github.com/cobbles/ant-simulator/server"
	"log"
)

func main() {
	srv := server.New()
	log.Println("Starting server on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
