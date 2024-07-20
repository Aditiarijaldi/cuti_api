package routes

import (
	"cuti_api/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/karyawan", controllers.GetKaryawans)
	e.GET("/karyawan/:id", controllers.GetKaryawan)
	e.POST("/karyawan", controllers.CreateKaryawan)
	e.PUT("/karyawan/:id", controllers.UpdateKaryawan)
	e.DELETE("/karyawan/:id", controllers.DeleteKaryawan)
}
