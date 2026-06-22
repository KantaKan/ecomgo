package main

import (
	"log"
	"os"
)


func main() {
	cfg := config {
		addr: ":8080",
		db: dbConfig{},

	}
	api := application{
		config: cfg,
	}

	if err := api.run(api.mount()); err != nil {
		log.Fatal("server has failed to start:", err)
		os.Exit(1)
		
		}
}