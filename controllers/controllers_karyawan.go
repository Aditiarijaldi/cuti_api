package controllers

import (
	"net/http"
	"strconv"

	"cuti_api/config"
	"cuti_api/models"

	"github.com/labstack/echo/v4"
)

// Get all karyawans
func GetKaryawans(c echo.Context) error {
	var karyawans []models.Karyawan
	if err := config.GetDB().Table("tbl_karyawan").Find(&karyawans).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Karyawan not found", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, karyawans)
}

// Get single karyawan
func GetKaryawan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID", "error": err.Error()})
	}
	var karyawan models.Karyawan
	if err := config.GetDB().Table("tbl_karyawan").First(&karyawan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Karyawan not found", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, karyawan)
}

// Create new karyawan
func CreateKaryawan(c echo.Context) error {
	karyawan := new(models.Karyawan)
	if err := c.Bind(karyawan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind data", "error": err.Error()})
	}
	if err := config.GetDB().Table("tbl_karyawan").Create(karyawan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create karyawan", "error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Karyawan created successfully", "data": karyawan})
}

// Update karyawan
func UpdateKaryawan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID", "error": err.Error()})
	}
	var karyawan models.Karyawan
	if err := config.GetDB().Table("tbl_karyawan").First(&karyawan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Karyawan not found", "error": err.Error()})
	}
	if err := c.Bind(&karyawan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Failed to bind data", "error": err.Error()})
	}
	if err := config.GetDB().Table("tbl_karyawan").Save(&karyawan).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update karyawan", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Karyawan updated successfully", "data": karyawan})
}

// Delete karyawan
func DeleteKaryawan(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid ID", "error": err.Error()})
	}
	if err := config.GetDB().Table("tbl_karyawan").Delete(&models.Karyawan{}, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Karyawan not found", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Karyawan deleted successfully"})
}
