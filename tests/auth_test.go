
package tests

import (
    "testing"
    "staff-contract-management/services"
    "staff-contract-management/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func TestAuthenticateUser(t *testing.T) {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    db.AutoMigrate(&models.Staff{})

    db.Create(&models.Staff{Username: "testuser", Password: "password", Role: "staff"})

    token, err := services.AuthenticateUser(db, "testuser", "password")
    if err != nil || token == "" {
        t.Errorf("Expected valid token, got error: %v", err)
    }
}
