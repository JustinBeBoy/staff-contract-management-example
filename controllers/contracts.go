package controllers

import (
	"net/http"
	"staff-contract-management/models"
	"staff-contract-management/utils"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func CreateContract(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		staff := user.(*models.Staff)

		var contract models.Contract
		if err := c.ShouldBindJSON(&contract); err != nil {
			utils.RespondError(c.Writer, http.StatusBadRequest, "Invalid request")
			return
		}

		contract.StaffID = staff.ID
		if err := db.Create(&contract).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusInternalServerError, "Could not create contract")
			return
		}

		utils.RespondJSON(c.Writer, http.StatusCreated, contract)
	}
}

func GetContracts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		staff := user.(*models.Staff)

		var contracts []models.Contract
		if staff.Role == "admin" {
			db.Find(&contracts)
		} else {
			db.Where("staff_id = ?", staff.ID).Find(&contracts)
		}
		utils.RespondJSON(c.Writer, http.StatusOK, contracts)
	}
}

func GetContractDetail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		staff := user.(*models.Staff)

		contractID := c.Param("id")
		var contract models.Contract
		if err := db.Where("id = ? AND (staff_id = ? OR ? = 'admin')", contractID, staff.ID, staff.Role).First(&contract).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusNotFound, "Contract not found")
			return
		}
		utils.RespondJSON(c.Writer, http.StatusOK, contract)
	}
}

func UpdateContract(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		staff := user.(*models.Staff)

		contractID := c.Param("id")
		var contract models.Contract
		if err := db.Where("id = ? AND (staff_id = ? OR ? = 'admin')", contractID, staff.ID, staff.Role).First(&contract).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusNotFound, "Contract not found")
			return
		}

		if err := c.ShouldBindJSON(&contract); err != nil {
			utils.RespondError(c.Writer, http.StatusBadRequest, "Invalid request")
			return
		}

		if err := db.Save(&contract).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusInternalServerError, "Could not update contract")
			return
		}

		utils.RespondJSON(c.Writer, http.StatusOK, contract)
	}
}

func DeleteContract(db *gorm.DB, nc *nats.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		staff := user.(*models.Staff)

		contractID := c.Param("id")
		var contract models.Contract
		if err := db.Where("id = ? AND (staff_id = ? OR ? = 'admin')", contractID, staff.ID, staff.Role).First(&contract).Error; err != nil {
			utils.RespondError(c.Writer, http.StatusNotFound, "Contract not found")
			return
		}

		// Publish to NATS for background task
		nc.Publish("delete.contract", []byte(contractID))

		// Respond immediately
		utils.RespondJSON(c.Writer, http.StatusNoContent, nil)
	}
}
