package controllers

import (
	"net/http"
	"staff-contract-management/services"
	"staff-contract-management/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.RespondError(c.Writer, http.StatusBadRequest, "Invalid request")
			return
		}

		token, err := services.AuthenticateUser(db, req.Username, req.Password)
		if err != nil {
			utils.RespondError(c.Writer, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		utils.RespondJSON(c.Writer, http.StatusOK, map[string]string{"token": token})
	}
}
