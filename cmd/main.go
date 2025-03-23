package main

import (
	"github.com/bluespada/timewise/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// TODO: we need to implement CLI command with viper, but for now just run the server
	server.RunApp()
}
