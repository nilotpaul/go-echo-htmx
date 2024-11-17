package main

import (
	"log"

	"github.com/nilotpaul/go-echo-htmx/api"
	"github.com/nilotpaul/go-echo-htmx/config"
)

func main() {
	// Loading the env vars from either a `.env` file or runtime.
	env := config.MustloadEnv()

	// API Server
	server := api.NewAPIServer(*env)
	log.Fatal(server.Start())
}