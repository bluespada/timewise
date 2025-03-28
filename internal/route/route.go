// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains routing for timewise including SPA route, Api Route, and GraphQL.
package route

import (
	"fmt"
	"time"

	timewise_config "github.com/bluespada/timewise/internal/config"
	"github.com/bluespada/timewise/internal/graph"
	"github.com/bluespada/timewise/internal/graph/playground"
	route_api "github.com/bluespada/timewise/internal/route/api"
	"github.com/bluespada/timewise/internal/utils/types"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {

	// initialize api route
	api := app.Group("/api")

	public_route := api.Group("/public")
	private_route := api.Group("/private")

	route_api.InitPublicApiRoute(public_route)
	route_api.InitPrivateRoute(private_route)

	// handle api information and global api route.
	api.All("/", handleApiInformation)
	api.All("*", handleApiNotFound)

	app.All("/graphql", graph.GraphHandler)
	app.Get("/graphql/playground", playground.HandlerGraphQLPlayground(graph.Schema))
}

func handleApiInformation(c *fiber.Ctx) error {
	res := types.NewApiResponse()
	res.Error = false
	res.Message = "Api Information"
	uptime := time.Since(timewise_config.APP_START_TIME)
	res.Data = map[string]interface{}{
		"name":           timewise_config.APP_NAME,
		"version":        timewise_config.APP_VERSION,
		"uptime":         fmt.Sprintf("%02dh:%02dm:%02ds", int(uptime.Hours()), int(uptime.Minutes())%60, int(uptime.Seconds())%60),
		"uptime_seconds": int(uptime.Seconds()),
	}

	return c.JSON(res)
}

func handleApiNotFound(c *fiber.Ctx) error {
	res := types.NewApiResponse()
	res.Error = true
	res.Message = "Endpoint not found."
	return c.Status(fiber.StatusNotFound).JSON(res)
}
