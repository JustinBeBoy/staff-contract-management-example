package controllers

import (
	"net/http"
	"staff-contract-management/models"
	"staff-contract-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func CreateStaff(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user role from context (this would typically be set by a middleware)
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.RespondError(c.Writer, http.StatusForbidden, "Only admin can create staff")
			return
		}

		var staff models.Staff
		if err := c.ShouldBindJSON(&staff); err != nil {
			utils.RespondError(c.Writer, http.StatusBadRequest, "Invalid request")
			return
		}

		if err := db.Create(&staff).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusInternalServerError, "Could not create staff")
			return
		}

		utils.RespondJSON(c.Writer, http.StatusCreated, staff)
	}
}

func GetStaffs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user role from context (this would typically be set by a middleware)
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.RespondError(c.Writer, http.StatusForbidden, "Only admin can get staff list")
			return
		}

		var staffs []models.Staff
		db.Find(&staffs)
		utils.RespondJSON(c.Writer, http.StatusOK, staffs)
	}
}

func GetStaffDetail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user role from context (this would typically be set by a middleware)
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.RespondError(c.Writer, http.StatusForbidden, "Only admin can get staff details")
			return
		}

		staffID := c.Param("id")
		var staff models.Staff
		if err := db.Where("id = ?", staffID).First(&staff).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusNotFound, "Staff not found")
			return
		}
		utils.RespondJSON(c.Writer, http.StatusOK, staff)
	}
}

func UpdateStaff(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user role from context (this would typically be set by a middleware)
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.RespondError(c.Writer, http.StatusForbidden, "Only admin can update staff")
			return
		}

		staffID := c.Param("id")
		var staff models.Staff
		if err := db.Where("id = ?", staffID).First(&staff).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusNotFound, "Staff not found")
			return
		}

		if err := c.ShouldBindJSON(&staff); err != nil {
			utils.RespondError(c.Writer, http.StatusBadRequest, "Invalid request")
			return
		}

		if err := db.Save(&staff).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusInternalServerError, "Could not update staff")
			return
		}

		utils.RespondJSON(c.Writer, http.StatusOK, staff)
	}
}

func DeleteStaff(db *gorm.DB, nc *nats.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user role from context (this would typically be set by a middleware)
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			utils.RespondError(c.Writer, http.StatusForbidden, "Only admin can delete staff")
			return
		}

		staffID := c.Param("id")

		// Publish to NATS for background task
		nc.Publish("delete.staff", []byte(staffID))

		// Respond immediately
		utils.RespondJSON(c.Writer, http.StatusNoContent, nil)
	}
}
