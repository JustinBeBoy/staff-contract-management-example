package routes

import (
	"staff-contract-management/controllers"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB, nc *nats.Conn) *gin.Engine {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login(db))
	}

	contract := r.Group("/contracts")
	{
		contract.POST("/", controllers.CreateContract(db))
		contract.GET("/", controllers.GetContracts(db))
		contract.GET("/:id", controllers.GetContractDetail(db))
		contract.PUT("/:id", controllers.UpdateContract(db))
		contract.DELETE("/:id", controllers.DeleteContract(db, nc))
	}

	staff := r.Group("/staffs")
	{
		staff.POST("/", controllers.CreateStaff(db))
		staff.GET("/", controllers.GetStaffs(db))
		staff.GET("/:id", controllers.GetStaffDetail(db))
		staff.PUT("/:id", controllers.UpdateStaff(db))
		staff.DELETE("/:id", controllers.DeleteStaff(db, nc))
	}

	return r
}
