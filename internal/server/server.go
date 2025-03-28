// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains http server code for timewise
package server

import (
	"fmt"
	"log"
	"os"

	"github.com/bluespada/timewise/internal/model"
	"github.com/bluespada/timewise/internal/route"
	"github.com/bluespada/timewise/internal/utils/database"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

// RunApp is the main entry point for the timewise server. It will start http server
// with gofiber and initialize all routes.
func RunApp() {
	// load dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	// initialize database
	database.Initialize()

	// register the model
	database.RegisterModels(
		&model.ModelAddress{},
		&model.ModelUsers{},
		&model.ModelAuth{},
		&model.ModelAuthLog{},
	)

	// initialize gofiber
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	// global middleware
	app.Use(logger.New(logger.ConfigDefault))

	// initialize route
	route.InitRoute(app)

	// http listen
	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))

}
