package controllers

import (
	"staff-contract-management/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", Login(db))
	}

	contract := r.Group("/contracts")
	{
		contract.POST("/", CreateContract(db))
		contract.GET("/", GetContracts(db))
		contract.GET("/:id", GetContractDetail(db))
		contract.PUT("/:id", UpdateContract(db))
		contract.DELETE("/:id", DeleteContract(db, utils.InitNATS()))
	}

	staff := r.Group("/staffs")
	{
		staff.POST("/", CreateStaff(db))
		staff.GET("/", GetStaffs(db))
		staff.GET("/:id", GetStaffDetail(db))
		staff.PUT("/:id", UpdateStaff(db))
		staff.DELETE("/:id", DeleteStaff(db, utils.InitNATS()))
	}
}
