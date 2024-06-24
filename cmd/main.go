package main

import (
	"staff-contract-management/routes"
	"staff-contract-management/tasks"
	"staff-contract-management/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	db := utils.InitDB()
	nc := tasks.InitNATS()

	// Start background task subscriptions
	go tasks.SubscribeTasks(nc, db)

	r := gin.Default()
	routes.SetupRouter(r, db, nc)
	r.Run(":8080")
}
