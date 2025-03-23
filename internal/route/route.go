package route

import (
	"fmt"
	"time"

	timewise_config "github.com/bluespada/timewise/internal/config"
	papi "github.com/bluespada/timewise/internal/route/api"
	"github.com/bluespada/timewise/internal/utils/types"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {

	// initialize api route
	api := app.Group("/api")
	papi.InitPublicApiRoute(api)

	// handle api information and global api route.
	api.All("/", handleApiInformation)
	api.All("*", handleApiNotFound)
}

func handleApiInformation(c *fiber.Ctx) error {
	res := types.NewApiResponse()
	res.Error = false
	res.Message = "Api Information"
	uptime := time.Since(timewise_config.APP_START_TIME)
	res.Data = map[string]interface{}{
		"name":        timewise_config.APP_NAME,
		"description": "Timewise API for Payroll and Attendance Management System.",
		"version":     timewise_config.APP_VERSION,
		"uptime":      fmt.Sprintf("%02dh:%02dm:%02ds", int(uptime.Hours()), int(uptime.Minutes())%60, int(uptime.Seconds())%60),
	}

	return c.JSON(res)
}

func handleApiNotFound(c *fiber.Ctx) error {
	res := types.NewApiResponse()
	res.Error = true
	res.Message = "Endpoint not found."
	return c.Status(fiber.StatusNotFound).JSON(res)
}
