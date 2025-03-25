// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit

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
