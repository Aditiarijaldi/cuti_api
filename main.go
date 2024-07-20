package main

import (
	"cuti_api/config"
	"cuti_api/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize database connection
	config.InitDB()

	// Route for root
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "API Success Berjalan")
	})

	// Routes for karyawan
	e.GET("/karyawan", controllers.GetKaryawans)
	e.GET("/karyawan/:id", controllers.GetKaryawan)
	e.POST("/karyawan", controllers.CreateKaryawan)
	e.PUT("/karyawan/:id", controllers.UpdateKaryawan)
	e.DELETE("/karyawan/:id", controllers.DeleteKaryawan)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
